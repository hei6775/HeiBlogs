package log

import (
	"fmt"
	l "log"
	"Algorithm/until/file"

	"github.com/astaxie/beego/logs"
)

type ILogger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Emergency(format string, args ...interface{})
	////////////////////////////////////
	SetLevel(int)
	Close()
}

func New(logDir string, logFileBaseName string) ILogger {
	file.MustMkdirIfNotExist(logDir)
	//chanl number
	var logger = logs.NewLogger(10000)
	//异步输出
	logger.Async()
	//输出文件名和行号
	logger.EnableFuncCallDepth(true)

	config := fmt.Sprintf(`{"filename":"%s/%s.log","level":%d,"maxlines":250000,"separate":["error"]}`,
		logDir, logFileBaseName, logs.LevelDebug)
	//mutifile配置
	logger.SetLogger("multifile", config)
	return logger
}

type ConsoleLogger struct {
}

//为logger实现接口
func (ConsoleLogger) Debug(format string, args ...interface{}) {
	l.Printf("[Debug] %s\n", fmt.Sprintf(format, args...))
}

func (ConsoleLogger) Info(format string, args ...interface{}) {
	l.Printf("[Info] %s\n", fmt.Sprintf(format, args...))
}

func (ConsoleLogger) Warn(format string, args ...interface{}) {
	l.Printf("[Warn] %s\n", fmt.Sprintf(format, args...))
}

func (ConsoleLogger) Error(format string, args ...interface{}) {
	l.Printf("[Error] %s\n", fmt.Sprintf(format, args...))
}

func (ConsoleLogger) Emergency(format string, args ...interface{}) {
	l.Printf("[Emergency] %s\n", fmt.Sprintf(format, args...))
}

func (ConsoleLogger) SetLevel(int) {

}

func (ConsoleLogger) Close() {

}
