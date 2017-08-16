package main

import (
	"fmt"
	"go_dev/day9/exercises/chat/chat_server/model"
	"net"
)

func runServer(addr string) (err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen failed,", err)
		return
	}


	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("listen failed,", err)
			continue
		}
		go proccess(conn)

	}

}

func proccess(conntion net.Conn) {
	defer conntion.Close()
	client := &model.Client{
		Conn: conntion,
	}
	err := client.Process()
	if err != nil {
		fmt.Println("client proccess failed", err)
		return
	}
}
