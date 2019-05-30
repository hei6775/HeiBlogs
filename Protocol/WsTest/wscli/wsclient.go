package main

import (
	"flag"
	"net/url"
	"github.com/gorilla/websocket"
	"fmt"
	"io/ioutil"
)

var serveraddr = flag.String("s1","127.0.0.1:9630","http service address")
var serverurl = url.URL{Scheme:"ws",Host:*serveraddr,Path:"/echo"}
var data = []byte(`bad new`)
func main(){
	client,_,_ := websocket.DefaultDialer.Dial(serverurl.String(),nil)
	client.WriteMessage(websocket.TextMessage,data)
	recMsgType, recMsg,err := client.ReadMessage()
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Println("recMsgType:",recMsgType)
	fmt.Println("recMsg:",string(recMsg))
	fmt.Println(client.LocalAddr().String())
	fmt.Println(client.LocalAddr().Network())

	_,iii,_ := client.NextReader()
	p2,_ :=ioutil.ReadAll(iii)
	fmt.Println("p2:",string(p2))

	client.Close()
}


