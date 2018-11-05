package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
)

//websocket连接，长连接


var upgrader = websocket.Upgrader{}
var data = []byte(`server send msg:good`)
func home(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"home.html")
}
func echo(w http.ResponseWriter,r *http.Request){
	con,err := upgrader.Upgrade(w,r,nil)
	if err !=nil {
		log.Printf("errors:[%v] \n",err)
		return
	}
	fmt.Println("server add:",con.LocalAddr().String())
	defer con.Close()
	for {
		recMsgType,recMsg,recMsgErr:=con.ReadMessage()
		if recMsgErr != nil {
			log.Printf("ReadMsg error: [%v] \n",recMsgErr)
			return
		}
		log.Printf("recMsgType: [%v] \n",recMsgType)
		log.Printf("recMsg: [%v] \n",string(recMsg))
		con.WriteMessage(websocket.TextMessage,data)

		con.WriteMessage(websocket.TextMessage,data)
	}
}

func main(){
	log.SetFlags(0)
	http.HandleFunc("/echo",echo)
	//http.HandleFunc("/",home)
	log.Fatal(http.ListenAndServe("localhost:9630", nil))
}