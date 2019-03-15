package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"time"
)

func Server(){
	pc,err := net.ListenPacket("udp","127.0.0.1:9631")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start Server with : %s ",pc.LocalAddr())
	defer pc.Close()

	clients := make([]net.Addr,0)
	go func() {
		for {
			for _,addr := range clients{
				_,err := pc.WriteTo([]byte(addr.String()),addr)
				if err != nil {
					log.Println(err)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		buf := make([]byte,256)
		n,addr,err := pc.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		clients = append(clients, addr)

		log.Println(string(buf[0:n]))
		log.Println(addr.String(), "connecting...", len(clients), "connected")
	}
}
func handleConnection(conn net.Conn){
	reader := bufio.NewReader(conn)
	for {
		data,_,err := reader.ReadLine()
		if err != nil{
			log.Println(err)
			return
		}
		fmt.Println("client:",string(data))
		_,err = conn.Write([]byte("word\n"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func main(){
	Server()
}