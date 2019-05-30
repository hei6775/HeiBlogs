package main

import (
	"log"
	"net"
	"time"
)

//socket连接

func main() {
	l, err := net.ListenPacket("udp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Start server with:%s", l.LocalAddr())
	defer l.Close()

	clients := make([]net.Addr, 0)
	go func() {
		for {
			for _, addr := range clients {
				_, err := l.WriteTo([]byte("pong\n"), addr)
				if err != nil {
					log.Println(err)
				}
			}

			time.Sleep(5 * time.Second)
		}
	}()

	for {
		buf := make([]byte, 256)
		n, addr, err := l.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		clients = append(clients, addr)

		log.Println(string(buf[0:n]))
		log.Println(addr.String(), "connecting...", len(clients), "connected")
	}
}
