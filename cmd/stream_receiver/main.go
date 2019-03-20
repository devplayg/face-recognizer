package main

import (
	"net"
	"fmt"
	"encoding/gob"
)

type Message struct {
	Text string
	Seq int
}

func handleConnection(conn net.Conn) {
	fmt.Print("New commer!")
	for {
		decoder := gob.NewDecoder(conn)
		msg := &Message{}
		if err := decoder.Decode(msg); err != nil {
			continue
		} else {
			fmt.Printf("msg: %s, seq: %d\n", msg.Text, msg.Seq);
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("Server is ready...")
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}