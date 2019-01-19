package logs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type fileLogWriter struct {
	Level int `json:"level"`
	sync.RWMutex
	//The opened file
	FileName   string `json:"filename"`
	fileWriter *os.File

	//Rotate at line
	MaxLines         int `json:"maxlines"`
	maxLinesCurLines int

	MaxSize        int `json:"maxsize"`
	maxSizeCurSize int

	Perm       string `json:"perm"`
	RotatePerm string `json:"rotateperm"`

	suffix string
}

func newFileWriter() Logger {
	f := &fileLogWriter{
		MaxLines: 10000,
		MaxSize:  1 << 28,
	}
	return f
}

func (f *fileLogWriter) Init(config string) error {
	err := json.Unmarshal([]byte(config), f)
	if err != nil {
		return err
	}
	if len(f.FileName) == 0 {
		return errors.New("config must have filename")
	}
	f.suffix = filepath.Ext(f.FileName)
	if f.suffix == "" {
		f.suffix = ".log"
	}
	err = f.startLogger()
	return err
}

func (f *fileLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	if level > f.Level {
		return nil
	}

	return nil
}

func (f *fileLogWriter) Destroy() {

}

func (f *fileLogWriter) Flush() {

}

func (f *fileLogWriter) startLogger() error {
	file, err := f.createLogFile()
	if err != nil {
		return nil
	}
	if f.fileWriter != nil {
		f.fileWriter.Close()
	}
	f.fileWriter = file
	return f.initFd()
}

func (f *fileLogWriter) createLogFile() (*os.File, error) {
	perm, err := strconv.ParseInt(f.Perm, 8, 64)
	if err != nil {
		return nil, err
	}
	filepath := path.Dir(f.FileName)
	os.MkdirAll(filepath, os.FileMode(perm))
	// 只写模式打开文件  写操作时将数据附加到文件尾部  如果不存在将创建一个新文件
	fd, err := os.OpenFile(f.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(perm))
	if err == nil {
		//修改指定文件对象的mode
		os.Chmod(f.FileName, os.FileMode(perm))
	}
	return fd, err
}

func (f *fileLogWriter) initFd() error {
	fd := f.fileWriter
	fInfo, err := fd.Stat()
	if err != nil {
		return fmt.Errorf("get stat err: %s", err)
	}
	f.maxSizeCurSize = int(fInfo.Size())
	f.maxLinesCurLines = 0

	if fInfo.Size() > 0 && f.MaxLines > 0 {
		count, err := f.lines()
		if err != nil {
			return err
		}
		f.maxLinesCurLines = count
	}
	return nil
}

func (f *fileLogWriter) lines() (int, error) {
	fd, err := os.Open(f.FileName)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	buf := make([]byte, 32768) //32k
	count := 0
	lineSep := []byte{'\n'}
	//统计行数
	for {
		c, err := fd.Read(buf)
		if err != nil && err != io.EOF {
			return count, err
		}
		count += bytes.Count(buf[:c], lineSep)
		if err == io.EOF {
			break
		}
	}
	return count, nil
}

//==========================================
func init() {
	Register(AdapterFile, newFileWriter)
}
