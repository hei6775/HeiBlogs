package logs

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestFilePerm(t *testing.T) {
	log := NewLogger(10000)
	// use 0666 as test perm cause the default umask is 022
	log.SetLogger("file", `{"filename":"test.log", "perm": "0666"}`)
	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warn("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	file, err := os.Stat("test.log")
	if err != nil {
		t.Fatal(err)
	}
	if file.Mode() != 0666 {
		t.Fatal("unexpected log file permission")
	}
	//os.Remove("test.log")
}

func TestFile1(t *testing.T) {
	log := NewLogger(10000)
	log.SetLogger("file", `{"filename":"test.log"}`)
	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warn("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	f, err := os.Open("test.log")
	if err != nil {
		t.Fatal(err)
	}
	b := bufio.NewReader(f)
	lineNum := 0
	for {
		line, _, err := b.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			lineNum++
		}
	}
	var expected = LevelDebug + 1
	if lineNum != expected {
		t.Fatal(lineNum, "not "+strconv.Itoa(expected)+" lines")
	}
	//os.Remove("test.log")
}
