package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func initServer(HostAndPort string)(*net.TCPListener){
	addr, err := net.ResolveTCPAddr("tcp",HostAndPort)
	checkerr(err)
	listener, err := net.ListenTCP("tcp",addr)
	checkerr(err)
	fmt.Println("Listen to: ",listener.Addr().String())
	return listener
}

func watchAndHandle(conn net.Conn){

	connForm := conn.RemoteAddr().String()
	fmt.Println("Connect from: ",connForm)
	connLocalAdr := conn.LocalAddr().String()
	fmt.Println("Connect Local Address: ",connLocalAdr)
	conn.Write([]byte(`Hello,I am server.`))

	for {
		inputReader := bufio.NewReader(os.Stdin)
		s2cmsg,_ := inputReader.ReadString('\n')
		trimmeds2cmsg := strings.Trim(s2cmsg,"\r\n")
		_,err := conn.Write([]byte("[server] says: "+trimmeds2cmsg))
		checkerr(err)
	}
}

func showmsg(conn net.Conn){
	for {
		buf := make([]byte,512)
		buflen,err := conn.Read(buf)
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(string(buf[0:buflen]))
	}
}

func main(){
	var Addr = "localhost:50000"
	fmt.Println("Server is working.")
	listener := initServer(Addr)
	for {
		conn, err := listener.Accept()
		checkerr(err)
		go watchAndHandle(conn)
		go showmsg(conn)
	}


}

func checkerr(err error){
	if err != nil {
		fmt.Printf("This error's Type:%T \n.This error :%v",err,err)
	}
	return
}