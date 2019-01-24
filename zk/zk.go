package zk

import "github.com/samuel/go-zookeeper/zk"

type ZK struct {
	Conn    *zk.Conn
	Servers []string
}
