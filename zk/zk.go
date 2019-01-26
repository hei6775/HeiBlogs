package zk

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type ZK struct {
	Conn      *zk.Conn
	Servers   []string `json:"servers"`
	Timeout   int      `josn:"timeout"`
	MsgChan   chan *RetChan
	functions map[interface{}]interface{}
}

type RetChan struct {
	id    interface{}
	etype int
	args  []interface{} //参数
	cb    interface{}   //回调函数
}

func NewGoZk(config string) (*ZK, error) {

	gozk := new(ZK)
	err := json.Unmarshal([]byte(config), gozk)
	if err != nil {
		return nil, fmt.Errorf("config [%v]", err)
	}
	if gozk.Timeout > 0 {
		gozk.Conn, _, err = zk.Connect(gozk.Servers, time.Duration(gozk.Timeout)*time.Second)
	} else {
		gozk.Conn, _, err = zk.Connect(gozk.Servers, 10*time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("Connnect invalid [%v]", err)
	}
	return gozk, nil
}

func (this *ZK) Register(id interface{}, f interface{}) {
	switch f.(type) {
	case func([]interface{}):
	case func([]interface{}) interface{}:
	case func([]interface{}) []interface{}:
	default:
		panic(fmt.Sprintf("function id %v:definition of function is invalid", id))
	}

	if _, ok := this.functions[id]; ok {
		panic(fmt.Sprintf("function is %v: already registered"))
	}

	this.functions[id] = f
}

func (this *ZK) execute(ci *RetChan) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("catch panic %v", r)
		}
	}()

	f := this.functions[ci.id]
	if f == nil {
		err = fmt.Errorf("function id %v: function not registered", ci.id)
		return
	}

	switch f.(type) {
	case func([]interface{}):
		f.(func([]interface{}))(ci.args)
	case func([]interface{}) interface{}:
		ret := f.(func([]interface{}) interface{})(ci.args)
		fmt.Println(ret)
	case func([]interface{}) []interface{}:
		ret := f.(func([]interface{}) []interface{})(ci.args)
		fmt.Println(ret)
	}
	return nil
}

func (this *ZK) sendChan(ci *RetChan) {
	this.MsgChan <- ci
}

func (this *ZK) setWatch() {
	for _, v := range this.functions {
		path := v.(string)

	}
}
