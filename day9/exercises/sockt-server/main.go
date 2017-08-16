package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read failed, err:", err)
			return
		}
		fmt.Println("read data:", string(buf))

	}

}

func server() {
	fmt.Println("start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("acceopt failed, err:", err)
			return
		}
		go process(conn)
	}

}

func main() {
	server()
}
