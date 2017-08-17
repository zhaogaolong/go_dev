package main

import (
	"fmt"
	"net"
)

var userId int
var passwd string

func main() {

	var (
		userId int
		passwd string
	)
	fmt.Scanf("%d %s\n", &userId, &passwd)
	// fmt.Println(userId, passwd)
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("Error Dial.....")
		return
	}
	err = login(conn, userId, passwd)

	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}
	// outputUserOnline()
	go proccessServerMsg(conn)
	for {
		logic()
	}
}
