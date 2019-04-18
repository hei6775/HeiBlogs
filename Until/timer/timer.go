package timer

import (
	"runtime"
	"time"

	"github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
)

// one dispatcher per goroutine (goroutine not safe)
//leaf的Timer的调度器
type Dispatcher struct {
	ChanTimer chan *Timer
}

//创建Dispatcher
func NewDispatcher(l int) *Dispatcher {
	disp := new(Dispatcher)
	disp.ChanTimer = make(chan *Timer, l)
	return disp
}

//Timer结构
type Timer struct {
	t  *time.Timer
	cb func()
}

func (t *Timer) Stop() {
	t.t.Stop()
	t.cb = nil
}

//执行回调函数，如果panic则捕获
//并且调用runtime.Stac来获取调用它的 goroutine 的格式化堆栈跟踪数据
//最后打印到日志
func (t *Timer) Cb() {
	defer func() {
		t.cb = nil
		if r := recover(); r != nil {
			if conf.LenStackBuf > 0 {
				buf := make([]byte, conf.LenStackBuf)
				l := runtime.Stack(buf, false)
				log.Error("%v: %s", r, buf[:l])
			} else {
				log.Error("%v", r)
			}
		}
	}()

	if t.cb != nil {
		t.cb()
	}
}

//Go 语言标准库提供了定时器的支持：
// AfterFunc 会等待 d 时长后调用 f 函数，这里的 f 函数将在另外一个 goroutine 中执行。
//leaf通过调用golang原生的定时器，然后这个定时器，会发送消息到Dispatcher的通道中
func (disp *Dispatcher) AfterFunc(d time.Duration, cb func()) *Timer {
	t := new(Timer)
	t.cb = cb
	t.t = time.AfterFunc(d, func() {
		disp.ChanTimer <- t
	})
	return t
}

// Cron
type Cron struct {
	t *Timer
}

func (c *Cron) Stop() {
	if c.t != nil {
		c.t.Stop()
	}
}

func (disp *Dispatcher) CronFunc(cronExpr *CronExpr, _cb func()) *Cron {
	c := new(Cron)

	now := time.Now()
	nextTime := cronExpr.Next(now)
	if nextTime.IsZero() {
		return c
	}

	// callback
	var cb func()
	cb = func() {
		defer _cb()

		now := time.Now()
		nextTime := cronExpr.Next(now)
		if nextTime.IsZero() {
			return
		}
		c.t = disp.AfterFunc(nextTime.Sub(now), cb)
	}

	c.t = disp.AfterFunc(nextTime.Sub(now), cb)
	return c
}
