package g

import (
	"container/list"
	"sync"
	"runtime"
)

type Go struct {
	ChanCb chan func()
	pendingGo int
}

type LinearGo struct {
	f func()  //执行函数
	cb func() //回调函数
}

//上下文
type LinearContext struct {
	g *Go
	linearGo *list.List //双向链表
	mutexLinearGo sync.Mutex
	mutexExecution sync.Mutex
}

func New(l int) *Go {
	g := new(Go)
	//缓存为1的函数类型通道
	g.ChanCb = make(chan func(),1)
	return g
}

func (g *Go)Go(f func(),cb func()){
	g.pendingGo++
	go func() {
		defer func() {
			g.ChanCb <- cb
			if r := recover();r!= nil {
				buf := make([]byte,4096)
				l := runtime.Stack(buf,false)

			}

		}()

		f()

	}()
}

