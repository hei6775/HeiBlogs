package logs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Name for adapter with beego official support
const (
	AdapterConsole   = "console"
	AdapterFile      = "file"
	AdapterMultiFile = "multifile"
)

// RFC5424 log message levels.
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

var adapters = make(map[string]newLoggerFunc)
var logMsgPool *sync.Pool

type newLoggerFunc func() Logger

func Register(name string, log newLoggerFunc) {
	if log == nil {
		panic("logs: Register provide is nil")
	}
	if _, dup := adapters[name]; dup {
		panic("log:Register called twice for provider " + name)
	}
	adapters[name] = log

}

type BeeLogger struct {
	lock                sync.Mutex
	level               int
	init                bool
	enableFuncCallDepth bool
	loggerFuncCallDepth int
	asynchronous        bool
	prefix              string
	msgChanLen          int64
	msgChan             chan *logMsg
	signalChan          chan string
	wg                  sync.WaitGroup
	outputs             []*nameLogger
}

const (
	defaultAsyncMsgLen = 1e3
)

type nameLogger struct {
	Logger
	name string
}

type logMsg struct {
	level int
	msg   string
	when  time.Time
}
type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

func NewLogger(channelLens ...int64) *BeeLogger {
	bl := new(BeeLogger)
	bl.level = LevelDebug
	//0-3 use 2
	bl.loggerFuncCallDepth = 2
	bl.msgChanLen = append(channelLens, 0)[0]
	if bl.msgChanLen <= 0 {
		bl.msgChanLen = defaultAsyncMsgLen
	}
	bl.signalChan = make(chan string, 1)
	bl.setLogger(AdapterConsole)
	return bl
}

func (bl *BeeLogger) setLogger(adapterName string, configs ...string) error {
	//avoid configs not exists
	config := append(configs, "{}")[0]
	for _, l := range bl.outputs {
		if l.name == adapterName {
			return fmt.Errorf("logs:duplicate adaptername %q (you have set this logger before)", adapterName)
		}
	}
	//get new log function
	log, ok := adapters[adapterName]
	if !ok {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}

	lg := log()
	err := lg.Init(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs.BeeLogger.SetLogger:"+err.Error())
		return err
	}
	bl.outputs = append(bl.outputs, &nameLogger{name: adapterName, Logger: lg})
	return nil
}

//SetLogger
func (bl *BeeLogger) SetLogger(adapterName string, configs ...string) error {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	if !bl.init {
		bl.outputs = []*nameLogger{}
		bl.init = true
	}
	return bl.setLogger(adapterName, configs...)
}

func (bl *BeeLogger) delLogger(adapterName string) error {
	outputs := []*nameLogger{}
	for _, lg := range bl.outputs {
		if lg.name == adapterName {
			lg.Logger.Destroy()
		} else {
			outputs = append(outputs, lg)
		}
	}

	if len(bl.outputs) == len(outputs) {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}
	bl.outputs = outputs
	return nil
}

//DelLogger
func (bl *BeeLogger) DelLogger(adapterName string) error {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	return bl.delLogger(adapterName)
}

// EnableFuncCallDepth enable log funcCallDepth
func (bl *BeeLogger) EnableFuncCallDepth(b bool) {
	bl.enableFuncCallDepth = b
}

//Async set the log to asynchronous and start the goroutine
func (bl *BeeLogger) Async(msgLen ...int64) *BeeLogger {
	bl.lock.Lock()
	defer bl.lock.Unlock()

	if bl.asynchronous {
		return bl
	}

	bl.asynchronous = true

	if len(msgLen) > 0 && msgLen[0] > 0 {
		bl.msgChanLen = msgLen[0]
	}

	bl.msgChan = make(chan *logMsg, bl.msgChanLen)
	logMsgPool = &sync.Pool{
		New: func() interface{} {
			return &logMsg{}
		},
	}
	bl.wg.Add(1)
	go bl.startLogger()
	return bl
}

//Write msg
func (bl *BeeLogger) writeToLoggers(when time.Time, msg string, level int) {
	for _, lg := range bl.outputs {
		err := lg.WriteMsg(when, msg, level)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to WriteMsg to adapter:%v,error:%v\n", lg.name, err)
		}
	}
}

func (bl *BeeLogger) writeMsg(logLevel int, msg string, v ...interface{}) error {
	if !bl.init {
		bl.lock.Lock()
		bl.setLogger(AdapterConsole)
		bl.lock.Unlock()
	}
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}

	msg = bl.prefix + " " + msg
	when := time.Now()
	//Caller
	if bl.enableFuncCallDepth {
		_, file, line, ok := runtime.Caller(bl.loggerFuncCallDepth)
		if !ok {
			file = "???"
			line = 0
		}
		_, filename := path.Split(file)
		msg = "[" + filename + ":" + strconv.Itoa(line) + "]" + msg
	}
	//asyc use pool
	if bl.asynchronous {
		lm := logMsgPool.Get().(*logMsg)
		lm.level = logLevel
		lm.msg = msg
		lm.when = when
		bl.msgChan <- lm
	} else {
		bl.writeToLoggers(when, msg, logLevel)
	}
	return nil
}

func (bl *BeeLogger) startLogger() {
	gameOver := false
	for {
		select {
		case bm := <-bl.msgChan:
			bl.writeToLoggers(bm.when, bm.msg, bm.level)
		case sg := <-bl.signalChan:
			bl.flush()
			if sg == "close" {
				for _, l := range bl.outputs {
					l.Destroy()
				}
				bl.outputs = nil
				gameOver = true
			}
			bl.wg.Done()
		}
		if gameOver {
			break
		}
	}
}

func (bl *BeeLogger) Close() {
	if bl.asynchronous {
		bl.signalChan <- "close"
		bl.wg.Wait()
		close(bl.msgChan)
	} else {
		bl.flush()
		for _, l := range bl.outputs {
			l.Destroy()
		}
		bl.outputs = nil
	}
	close(bl.signalChan)
}

// Flush flush all chan data.
func (bl *BeeLogger) Flush() {
	if bl.asynchronous {
		bl.signalChan <- "flush"
		bl.wg.Wait()
		bl.wg.Add(1)
		return
	}
	bl.flush()
}

func (bl *BeeLogger) flush() {
	if bl.asynchronous {
		for {
			if len(bl.msgChan) > 0 {
				bm := <-bl.msgChan
				bl.writeToLoggers(bm.when, bm.msg, bm.level)
				logMsgPool.Put(bm)
				continue
			}
			break
		}
	}
	for _, l := range bl.outputs {
		l.Flush()
	}
}

////////////////////////////////////////////
func (bl *BeeLogger) Emergency(format string, v ...interface{}) {
	if LevelEmergency > bl.level {
		return
	}
	format = "[Emerg] " + format
	bl.writeMsg(LevelEmergency, format, v...)
}

func (bl *BeeLogger) Alert(format string, v ...interface{}) {
	if LevelAlert > bl.level {
		return
	}
	format = "[Alert] " + format
	bl.writeMsg(LevelAlert, format, v...)
}

func (bl *BeeLogger) Critical(format string, v ...interface{}) {
	if LevelCritical > bl.level {
		return
	}
	format = "[Criti] " + format
	bl.writeMsg(LevelCritical, format, v...)
}

func (bl *BeeLogger) Error(format string, v ...interface{}) {
	if LevelError > bl.level {
		return
	}
	format = "[Error] " + format
	bl.writeMsg(LevelError, format, v...)
}

func (bl *BeeLogger) Warn(format string, v ...interface{}) {
	if LevelWarning > bl.level {
		return
	}
	format = "[Warn] " + format
	bl.writeMsg(LevelWarning, format, v...)
}

func (bl *BeeLogger) Notice(format string, v ...interface{}) {
	if LevelNotice > bl.level {
		return
	}
	format = "[Notice] " + format
	bl.writeMsg(LevelNotice, format, v...)
}

func (bl *BeeLogger) Info(format string, v ...interface{}) {
	if LevelInformational > bl.level {
		return
	}
	format = "[Info] " + format
	bl.writeMsg(LevelInformational, format, v...)
}

func (bl *BeeLogger) Debug(format string, v ...interface{}) {
	if LevelDebug > bl.level {
		return
	}
	format = "[Debug] " + format
	bl.writeMsg(LevelDebug, format, v...)
}

//============================
// beeLogger references the used application logger.
var beeLogger = NewLogger()
