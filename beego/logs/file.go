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

	fileNameOnly, suffix string
}

func newFileWriter() Logger {
	f := &fileLogWriter{
		MaxLines: 10000,
		MaxSize:  1 << 28,
		Perm:     "0660",
		Level:    LevelDebug,
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
	tstr, _, _ := formatTimeHeader(when)
	msg = string(tstr) + msg + "\n"
	if f.isRotate() {
		f.Lock()
		f.doRotate(when)
		f.Unlock()
	}

	f.Lock()
	_, err := f.fileWriter.Write([]byte(msg))
	if err == nil {
		f.maxLinesCurLines++
		f.maxSizeCurSize += len(msg)
	}
	f.Unlock()
	return err
}

func (f *fileLogWriter) Destroy() {
	f.fileWriter.Close()
}

//from memory to disk
func (f *fileLogWriter) Flush() {
	f.fileWriter.Sync()
}

func (f *fileLogWriter) startLogger() error {
	file, err := f.createLogFile()
	if err != nil {
		return err
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

func (f *fileLogWriter) isRotate() bool {
	return (f.MaxSize > 0 && f.maxSizeCurSize > f.MaxSize) ||
		(f.MaxLines > 0 && f.maxLinesCurLines > f.MaxLines)
}

//create new file
func (f *fileLogWriter) doRotate(logtime time.Time) error {
	fName := ""

	rotatePerm, err := strconv.ParseInt(f.RotatePerm, 8, 64)
	if err != nil {
		return err
	}
	//获取文件信息
	_, err = os.Lstat(f.FileName)
	if err != nil {

	}
	format := "2006010215"

	if f.MaxLines > 0 || f.MaxSize > 0 {
		fName = f.fileNameOnly + fmt.Sprintf(".%s.%s", logtime.Format(format), f.suffix)
		_, err = os.Lstat(fName)
	} else {
		fName = f.fileNameOnly + fmt.Sprintf(".%s.%s", logtime.Format(format), f.suffix)
		_, err = os.Lstat(fName)
	}
	//文件已经存在
	if err == nil {
		return fmt.Errorf("doRotate cannot find rename:%s", f.FileName)
	}

	f.fileWriter.Close()

	//rename the file
	err = os.Rename(f.FileName, fName)
	if err != nil {
		goto ReStart_Log
	}
	//修改文件的mode
	err = os.Chmod(fName, os.FileMode(rotatePerm))

ReStart_Log:
	startLoggerErr := f.startLogger()

	if startLoggerErr != nil {
		return fmt.Errorf("Rotate startLogger: %s", startLoggerErr)
	}
	if err != nil {
		return fmt.Errorf("Rotate: %s", err)
	}
	return nil
}

//==========================================
func init() {
	Register(AdapterFile, newFileWriter)
}
