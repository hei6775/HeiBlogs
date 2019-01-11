package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"time"
)

var conurl = "127.0.0.1:9631"

var contype = "tcp"
//
//func Dial()(){
//	conn,err :=net.Dial(contype,conurl)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	defer conn.Close()
//
//	_,err = fmt.Fprintf(conn,"GET / HTTP/1.0\r\n\r\n")
//	if err != nil {
//		log.Fatal("err")
//	}
//
//	data,err := ioutil.ReadAll(conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(data))
//}

func Dial2(){
	conn, err := net.Dial(contype,conurl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		_,err := conn.Write([]byte("hello\n"))
		if err != nil{
			log.Fatal(err)
		}
		data,_,err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("server:",string(data))
		time.Sleep(5 * time.Second)
	}
}

func main(){
	Dial2()
}

