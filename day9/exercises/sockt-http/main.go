package main

import (
	"fmt"
	"io"
	"net"
)

func getHTTP() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error dialing...", err)
		return
	}

	io.Writer
	defer conn.Close()

	msg := "GET / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection:close\r\n"
	msg += "\r\n\r\n"

	_, err = io.WriteString(conn, msg)
	// _, err = conn.Write([]byte(msg))

	if err != nil {
		fmt.Println("write string failed, ", err)
		return
	}

	buf := make([]byte, 4096)

	for {
		count, err := conn.Read(buf)
		fmt.Println(string(count))
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}

}

func main() {
	getHTTP()
}
