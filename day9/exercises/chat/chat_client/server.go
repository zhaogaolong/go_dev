package main

import (
	"encoding/json"
	"fmt"
	"go_dev/day9/exercises/chat/proto"
	"net"
	"os"
)

func proccessServerMsg(conn net.Conn) {
	msg, err := readPackage(conn)
	if err != nil {
		fmt.Println(" read package err:", err)
		os.Exit(0)
	}
	var userStatus proto.UserStatusNotify
	err = json.Unmarshal([]byte(msg.Data), &userStatus)
	if err != nil {
		fmt.Println("user status notify data unmarshal err:", err)
		return
	}
	switch msg.Cmd {
	case proto.UserStatusNotifyRes:
		updateUserStatus(userStatus)
	}
}
