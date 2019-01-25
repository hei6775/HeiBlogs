package zk

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"testing"
	"time"
)

var hosts = []string{"192.168.159.128"}

func TestZK(t *testing.T) {
	conn, _, _ := zk.Connect(hosts, 5*time.Second)
	defer conn.Close()
	data1, sta, err1 := conn.Get("/root")
	fmt.Println(string(data1), sta, err1)
	//if err != nil {
	//	fmt.Println("===", err)
	//}
	//data1, sta, err1 := conn.Get("/test")
	//fmt.Println(string(data1), sta, err1)
	//var path = "/test"
	//var data = []byte("hello zk")
	//var flags = 0
	////flags有4种取值：
	////0:永久，除非手动删除
	////zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	////zk.FlagSequence  = 2:会自动在节点后面添加序号
	////3:Ephemeral和Sequence，即，短暂且自动添加序号
	//var acls = zk.WorldACL(zk.PermAll) //控制访问权限模式
	//
	//p, err_create := conn.Create(path, data, int32(flags), acls)
	//fmt.Println(p, err_create)
}
