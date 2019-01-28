package zk

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

const (
	TypeCreateNodeLasting     = iota //持久节点
	TypeCreateNodeEphemeral          //暂态节点
	TypeCreateNodeLasSequence        //持久有序节点
	TypeCreateNodeEphSequence        //暂态有序节点
	TypeDeleteNode                   //删除节点4
	TypeSetNodeDate                  //修改节点数据5
	TypeGetNodeData                  //获取节点数据6
	TypeExistNode                    //节点存在与否7
	TypeGetChilds                    //节点存在与否7
)

const (
	PermRead = 1 << iota
	PermWrite
	PermCreate
	PermDelete
	PermAdmin
	PermAll = 0x1f
)

type ZK struct {
	Conn      *zk.Conn
	Servers   []string `json:"servers"`
	Timeout   int      `josn:"timeout"`
	MsgChan   chan *MsgChan
	eChan     <-chan zk.Event
	functions map[interface{}]Funcs
}

type Funcs interface {
	DealWith([]byte, string) error
	OnDestry()
}

type MsgChan struct {
	Type interface{}
	Id   interface{}
	Args []interface{} //参数
	Cb   interface{}   //回调函数
}

//
type CallArg struct {
	Id    interface{}
	State interface{}
	Data  interface{}
	Perm  interface{}
	Error error
}

func NewGoZk(config string) (*ZK, error) {

	gozk := new(ZK)
	err := json.Unmarshal([]byte(config), gozk)
	if err != nil {
		return nil, fmt.Errorf("config [%v]", err)
	}
	if gozk.Timeout > 0 {
		gozk.Conn, gozk.eChan, err = zk.Connect(gozk.Servers, time.Duration(gozk.Timeout)*time.Second)
	} else {
		gozk.Conn, gozk.eChan, err = zk.Connect(gozk.Servers, 10*time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("Connnect invalid [%v]", err)
	}
	return gozk, nil
}

func (this *ZK) Register(id interface{}, funcs Funcs) {
	if _, ok := this.functions[id]; ok {
		panic(fmt.Sprintf("Funcs interface is %v: already registered"))
	}

	this.functions[id] = funcs
}

func (this *ZK) Run() {
	for {
		select {
		case eve := <-this.eChan:
			if eve.Type == zk.EventNodeDataChanged {
				this.resetWatch(eve.Path)
			}
			if eve.State == zk.StateDisconnected {
				this.stop()
				break
			}
		}
	}
}

func (this *ZK) sendChan(ci *MsgChan) {
	this.MsgChan <- ci
}

func (this *ZK) setWatch() {
	for k, f := range this.functions {
		path := k.(string)
		data, _, _, err := this.Conn.GetW(path)
		if err != nil {
			panic(fmt.Sprintf("path[%v] err[%v]", path, err))
		}

		if f == nil {
			err = fmt.Errorf("function id %v: function not registered")
			return
		}
		err = f.DealWith(data, path)
		panic(err)
	}
}

func (this *ZK) resetWatch(path string) error {
	data, _, _, err := this.Conn.GetW(path)
	if err != nil {
		return err
	}

	f := this.functions[path]
	if f == nil {
		err = fmt.Errorf("function id %v: function not registered", path)
		return err
	}
	err = f.DealWith(data, path)
	return err

}

func (this *ZK) stop() {

}

func (this *ZK) execute(ci *MsgChan) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("catch panic %v", r)
		}
	}()

	switch ci.Type {
	case TypeCreateNodeLasting:
	case TypeCreateNodeEphemeral:
	case TypeCreateNodeLasSequence:
	case TypeCreateNodeEphSequence:
	case TypeDeleteNode:
	case TypeSetNodeDate:
	case TypeGetNodeData:
	case TypeExistNode:
	default:
		err = fmt.Errorf("invalid ci type: %v", ci.Type)
		return err
	}
	callArg := new(CallArg)
	args := ci.Args
	tarArgs := []interface{}{
		callArg,
	}
	tarArgs = append(tarArgs, args...)
	this.callBack(ci.Cb, tarArgs)

	return nil
}

func (this *ZK) callBack(cb interface{}, args []interface{}) {
	//callBack
	switch cb.(type) {
	case func([]interface{}):
		cb.(func([]interface{}))(args)
	case func([]interface{}) interface{}:
		_ = cb.(func([]interface{}) interface{})(args)
	case func([]interface{}) []interface{}:
		_ = cb.(func([]interface{}) []interface{})(args)
	}
	return
}

func (this *ZK) createNodes(data, path, flag, perm interface{}) (outpath string, err error) {
	tarpath := path.(string)
	tardata := data.([]byte)
	tarflag := flag.(int32)
	var tarperm = zk.WorldACL(zk.PermAll)
	outpath, err = this.Conn.Create(tarpath, tardata, tarflag, tarperm)
	return
}

func (this *ZK) deleteNode(path interface{}) error {
	tarpath := path.(string)
	err := this.Conn.Delete(tarpath, -1)
	return err
}

func (this *ZK) existNode(data, path interface{}) (exist bool, err error) {
	tarpath := path.(string)
	exist, _, err = this.Conn.Exists(tarpath)
	return
}

func (this *ZK) getNodeData(path interface{}) (data []byte, err error) {
	tarpath := path.(string)
	data, _, err = this.Conn.Get(tarpath)
	return
}

func (this *ZK) setNodeData(data, path interface{}) error {
	tardata := data.([]byte)
	tarpath := path.(string)
	_, err := this.Conn.Set(tarpath, tardata, -1)
	return err
}

func (this *ZK) getchild(path interface{}) ([]string, error) {
	tarpath := path.(string)
	childrens, _, err := this.Conn.Children(tarpath)
	return childrens, err
}
