package Zk

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
	TypeUpdateNodeDate               //修改节点数据5
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

//封装zk结构
type ZK struct {
	Conn      *zk.Conn
	Servers   []string `json:"servers"`
	Timeout   int      `josn:"timeout"`
	MsgChan   chan *MsgChan
	eChan     <-chan zk.Event
	functions map[interface{}]Funcs
}

//函数接口
type Funcs interface {
	Done([]byte, string) error
	OnDestry()
}

//通道信息
type MsgChan struct {
	Type interface{}   //消息类型
	Id   interface{}   //ID
	Data interface{}   //数据
	Perm interface{}   //权限
	Args []interface{} //参数
	Cb   interface{}   //回调函数
}

////
//type CallArg struct {
//	Id    interface{}
//	State interface{}
//	Data  interface{}
//	Perm  interface{}
//	Error error
//}

//初始化zk
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

//注册节点
func (this *ZK) Register(id interface{}, funcs Funcs) {
	if _, ok := this.functions[id]; ok {
		panic(fmt.Sprintf("Funcs interface is %v: already registered"))
	}
	this.functions[id] = funcs
}

//zk run
//zk 节点变化执行zk的resetWatch函数
func (this *ZK) Run() {
	for {
		select {
		case eve := <-this.eChan:
			if eve.Type == zk.EventNodeDataChanged {
				this.resetWatch(eve.Path)
			}
			//zk断掉
			if eve.State == zk.StateDisconnected {
				this.stop()
				break
			}
		}
	}
}

//zk通道发送信息
func (this *ZK) sendChan(ci *MsgChan) {
	this.MsgChan <- ci
}

//全部设置监听
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
		err = f.Done(data, path)
		if err != nil {
			panic(err)
		}
	}
}

//重新设置监听
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
	err = f.Done(data, path)
	return err

}

//stop函数
//执行ondestry函数
func (this *ZK) stop() {
	for _, v := range this.functions {
		v.OnDestry()
	}
}

//收到消息，执行execute函数
func (this *ZK) execute(ci *MsgChan) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("catch panic %v", r)
		}
	}()

	switch ci.Type {
	case TypeCreateNodeLasting /*创建持久化节点*/ :
		this.createNodes(ci.Data, ci.Id, TypeCreateNodeLasting, ci.Perm)
	case TypeCreateNodeEphemeral /*创建暂态节点*/ :
		this.createNodes(ci.Data, ci.Id, TypeCreateNodeEphemeral, ci.Perm)
	case TypeCreateNodeLasSequence /*创建持久有序节点*/ :
		this.createNodes(ci.Data, ci.Id, TypeCreateNodeLasSequence, ci.Perm)
	case TypeCreateNodeEphSequence /*创建暂态有序节点*/ :
		this.createNodes(ci.Data, ci.Id, TypeCreateNodeEphSequence, ci.Perm)
	case TypeDeleteNode /*删除节点*/ :
		err := this.deleteNode(ci.Id)
		if err != nil {
			return err
		}
		this.functions[ci.Id].OnDestry()
	case TypeUpdateNodeDate /*更新节点数据*/ :
		err := this.updateNodeData(ci.Data, ci.Id)
		if err != nil {
			return err
		}
	case TypeGetNodeData /*获取节点数据*/ :
		data, err := this.getNodeData(ci.Id)
		if err != nil {
			return err
		}
		ci.Args = append(ci.Args, data)
	case TypeExistNode /*节点是否存在*/ :
		isExist, err := this.existNode(ci.Id)
		if err != nil {
			return err
		}
		ci.Args = append(ci.Args, isExist)
	case TypeGetChilds /*获取子节点ID*/ :
		childs, err := this.getchild(ci.Id)
		if err != nil {
			return err
		}
		ci.Args = append(ci.Args, childs)
	default:
		err := fmt.Errorf("invalid ci type: %v", ci.Type)
		return err
	}
	if ci.Cb != nil {
		go this.callBack(ci.Cb, ci.Args)
	}

	return nil
}

//回调函数
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

//创建节点
func (this *ZK) createNodes(data, path, flag, perm interface{}) (outpath string, err error) {
	tarpath := path.(string)
	tardata := data.([]byte)
	tarflag := flag.(int32)
	var tarperm = zk.WorldACL(zk.PermAll)
	outpath, err = this.Conn.Create(tarpath, tardata, tarflag, tarperm)
	return
}

//删除节点
func (this *ZK) deleteNode(path interface{}) error {
	tarpath := path.(string)
	err := this.Conn.Delete(tarpath, -1)
	return err
}

//节点是否存在
func (this *ZK) existNode(path interface{}) (exist bool, err error) {
	tarpath := path.(string)
	exist, _, err = this.Conn.Exists(tarpath)
	return
}

//获取节点数据
func (this *ZK) getNodeData(path interface{}) (data []byte, err error) {
	tarpath := path.(string)
	data, _, err = this.Conn.Get(tarpath)
	return
}

//更新节点数据
func (this *ZK) updateNodeData(data, path interface{}) error {
	tardata := data.([]byte)
	tarpath := path.(string)
	_, err := this.Conn.Set(tarpath, tardata, -1)
	return err
}

//获取指定节点的子节点
func (this *ZK) getchild(path interface{}) ([]string, error) {
	tarpath := path.(string)
	childrens, _, err := this.Conn.Children(tarpath)
	return childrens, err
}
