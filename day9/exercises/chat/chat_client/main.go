package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day9/exercises/chat/proto"
	"net"
	"os"
	"time"
)

func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var buf [8192]byte
	n, err := conn.Read(buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	// fmt.Println("read Package :", buf[0:4])
	var packLen uint32
	packLen = binary.BigEndian.Uint32(buf[0:4])
	// fmt.Printf("receive len:%d\n", packLen)

	n, err = conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read boby failed")
		return
	}

	// fmt.Printf("receive data: %s\n", string(buf[0:packLen]))

	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("data unmarshal failed")
	}
	return
}

func register(conn net.Conn, UserID int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserRegister

	var registerCmd proto.RegisterCmd
	registerCmd.User.UserID = UserID
	registerCmd.User.Nick = "stu01"
	registerCmd.User.Sex = "man"
	registerCmd.User.Passwd = passwd
	registerCmd.User.Header = "http://ww.baidu.com/img"

	data, err := json.Marshal(registerCmd)
	if err != nil {
		fmt.Println("user register json marshal failed, err:", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("user register json marshal failed, err:", err)
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))
	fmt.Println("packlen", packLen)
	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conn.Write(buf[:])

	if err != nil || n != 4 {
		fmt.Println("conn write failed")
		return
	}

	// fmt.Println("register data:", string(data))
	_, err = conn.Write([]byte(data))

	if err != nil {
		return
	}

	msg, err = readPackage(conn)
	if err != nil {
		fmt.Println("read package failed")

	}

	fmt.Println(msg)
	return

}

func login(conn net.Conn, userId int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLogin

	var loginCmd proto.LoginCmd
	loginCmd.Id = userId
	loginCmd.Passwd = passwd
	data, err := json.Marshal(loginCmd)

	if err != nil {
		return
	}

	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))
	// fmt.Println("packlen", packLen)
	binary.BigEndian.PutUint32(buf[0:4], packLen)
	n, err := conn.Write(buf[:])

	if err != nil || n != 4 {
		fmt.Println("conn write failed")
		return
	}

	_, err = conn.Write([]byte(data))

	if err != nil {
		return
	}

	msg, err = readPackage(conn)
	if err != nil {
		fmt.Println("read package failed")

	}

	fmt.Println(msg)

	var loginRes proto.LoginCmdResult
	json.Unmarshal([]byte(msg.Data), &loginRes)
	if loginRes.Code == 500 {
		fmt.Println("user not fount, must register user...")
		register(conn, userId, passwd)
		os.Exit(0)
	}

	fmt.Println("online user list:")
	for _, v := range loginRes.User {
		if v == userId {
			continue
		}
		fmt.Printf("user: %d\n", v)
	}

	return

}

func main() {

	var (
		userId int
		passwd string
	)
	fmt.Scanf("%d %s", &userId, &passwd)
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
	time.Sleep(time.Second * 10)

}
