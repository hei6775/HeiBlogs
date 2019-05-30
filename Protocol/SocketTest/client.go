package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	_, err = conn.Write([]byte("ping..."))
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(conn)
	for {
		dat, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(dat))
	}
}
