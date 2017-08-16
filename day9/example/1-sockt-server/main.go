package main

import "fmt"
import "net"

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		fmt.Println("read:", string(buf))
	}
}

func main() {
	fmt.Println("Start server.....")
	listen, err := net.Listen("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Listen fail ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept error ", err)
			return
		}

		go process(conn)

	}

}
