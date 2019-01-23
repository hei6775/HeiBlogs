package logs

import "testing"

func testConsole(bl *BeeLogger) {
	bl.Emergency("emergency")
	bl.Alert("alter")
	bl.Critical("critical")
	bl.Error("error")
	bl.Warn("warning")
	bl.Notice("notice")
	bl.Info("informational")
	bl.Debug("debug")
}

func TestConsole(t *testing.T) {
	log1 := NewLogger(10000)
	log1.EnableFuncCallDepth(true)
	log1.SetLogger("console", "")
	testConsole(log1)

	log2 := NewLogger(100)
	log2.SetLogger("console", `{"level":3}`)
	testConsole(log2)

	//output
	//2019-01-22 19:46:27.3856002 +0800 CST m=+0.024000001[console_test.go:6] emergency
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001[console_test.go:7] alter
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001[console_test.go:8] critical
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001[console_test.go:9] error
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001[console_test.go:10] warning
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001[console_test.go:11] notice
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001[console_test.go:12] informational
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001[console_test.go:13] debug
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001 emergency
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001 alter
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001 critical
	//2019-01-22 19:46:27.4146002 +0800 CST m=+0.053000001 error
}
