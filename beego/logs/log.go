package logs

import (
	"sync"
	"time"
)

const (
	AdapterFile = "file"
)

// RFC5424 log message levels.
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInfomational
	LevelDebug
)

const (
	defaultAsyncMsgLen = 1e3
)

var adapters = make(map[string]newLoggerFunc)

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

func NewLogger(channelLens ...int64) *BeeLogger {
	bl := new(BeeLogger)
	bl.level = LevelDebug
	bl.loggerFuncCallDepth = 2
	bl.msgChanLen = append(channelLens, 0)[0]
	if bl.msgChanLen <= 0 {
		bl.msgChanLen = defaultAsyncMsgLen
	}
	bl.signalChan = make(chan string, 1)
	return bl
}
