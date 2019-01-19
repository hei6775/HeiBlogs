package logs

import (
	"encoding/json"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

type brush func(string) string

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m" //取消颜色设置
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colors = []brush{
	newBrush("1;37"), // Emergency          white
	newBrush("1;36"), // Alert              cyan
	newBrush("1;35"), // Critical           magenta
	newBrush("1;31"), // Error              red
	newBrush("1;33"), // Warning            yellow
	newBrush("1;32"), // Notice             green
	newBrush("1;34"), // Informational      blue
	newBrush("1;44"), // Debug              Background blue
}

type logWriter struct {
	sync.Mutex
	writer io.Writer
}

func newLogWriter(wr io.Writer) *logWriter {
	return &logWriter{writer: wr}
}

func (lg *logWriter) println(when time.Time, msg string) {
	lg.Lock()
	//time to string
	h := []byte(when.String())
	lg.writer.Write(append(append(h, msg...), '\n'))
	lg.Unlock()
}

type consoleWriter struct {
	lg       *logWriter
	Level    int  `json:"level"`
	Colorful bool `json:"colorful"` //this field is useful only when system's terminal supports color
}

func NewConsole() Logger {
	cw := &consoleWriter{
		lg:       newLogWriter(os.Stdout),
		Level:    LevelDebug,
		Colorful: runtime.GOOS != "windows",
	}
	return cw
}

func (cw *consoleWriter) Init(jsonConfig string) error {
	if len(jsonConfig) == 0 {
		return nil
	}
	err := json.Unmarshal([]byte(jsonConfig), cw)
	if runtime.GOOS == "windows" {
		cw.Colorful = false
	}
	return err
}

func (cw *consoleWriter) WriteMsg(when time.Time, msg string, level int) error {
	if level > cw.Level {
		return nil
	}
	if cw.Colorful {
		msg = colors[level](msg)
	}
	cw.lg.println(when, msg)
	return nil
}

func (cw *consoleWriter) Destroy() {

}

func (cw *consoleWriter) Flush() {

}

//==========================================
func init() {
	Register("AdapterConsole", NewConsole)
}
