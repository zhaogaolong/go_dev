package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Can't not connt baidu.com")
		return
	}

	defer conn.Close()

	msg := "GET / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection: close \r\n"
	msg += "\r\n\r\n"

	n, err := io.WriteString(conn, msg)

	if err != nil {
		fmt.Println("Write bybe faild to baidu.com")
		return
	}
	fmt.Print("Write to baidu.com", n)
	buf := make([]byte, 4096)

	for {
		count, err := conn.Read(buf)
		fmt.Print("conut:", count, "error", err)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))

	}

}
