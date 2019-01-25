package zk

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type ZK struct {
	Conn    *zk.Conn
	Servers []string `json:"servers"`
	timeout int      `josn:"timeout"`
}

func NewGoZk(config string) (*ZK, error) {

	gozk := new(ZK)
	err := json.Unmarshal([]byte(config), gozk)
	if err != nil {
		return nil, fmt.Errorf("config [%v]", err)
	}
	if gozk.timeout > 0 {
		gozk.Conn, _, err = zk.Connect(gozk.Servers, time.Duration(gozk.timeout)*time.Second)
	} else {
		gozk.Conn, _, err = zk.Connect(gozk.Servers, 10*time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("Connnect invalid [%v]", err)
	}
	return gozk, nil
}

func (gozk *ZK) Get() {

}
