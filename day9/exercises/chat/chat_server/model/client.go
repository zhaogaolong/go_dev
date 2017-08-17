package model

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"go_dev/day9/exercises/chat/proto"
)

type Client struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (p *Client) readPackage() (msg proto.Message, err error) {
	n, err := p.Conn.Read(p.Buf[0:4])

	if n != 4 {
		err = errors.New("read header failed")
		return
	}

	fmt.Println("read package:", n)

	var packLen uint32
	packLen = binary.BigEndian.Uint32(p.Buf[0:4])
	fmt.Println("receive len:%d", packLen)

	n, err = p.Conn.Read(p.Buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read boby failed")
		return
	}

	err = json.Unmarshal(p.Buf[0:packLen], &msg)
	// fmt.Printf("receive data:%s\n", string(p.Buf[0:packLen]))
	if err != nil {
		fmt.Println("unmarshal failed, err", err)
	}
	return

}

func (p *Client) writePackage(data []byte) (err error) {
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(p.Buf[0:4], packLen)
	n, err := p.Conn.Write(p.Buf[0:4])
	if err != nil {
		fmt.Println("write header data failed")
		return
	}

	n, err = p.Conn.Write(data)
	if err != nil {
		fmt.Println("write data failed")
		return
	}
	if n != int(packLen) {
		fmt.Println("write data not finshed")
		err = errors.New("write data not finshed")
		return
	}
	return

}

func (p *Client) Process() (err error) {
	for {
		var msg proto.Message

		msg, err = p.readPackage()
		if err != nil {
			return
		}
		err = p.processMgr(msg)
		if err != nil {
			fmt.Println("process msg failed, err", err)
			continue
		}
	}
}

func (p *Client) processMgr(msg proto.Message) (err error) {
	fmt.Println("processMgr:", msg.Cmd)
	switch msg.Cmd {
	case proto.UserLogin:
		p.login(msg)
	case proto.UserRegister:
		p.register(msg)
	default:
		err = errors.New("unsupport message")
		return
	}
	return
}

func (p *Client) loginResp(err error) {
	var resMsg proto.Message
	resMsg.Cmd = proto.UserLoginResult

	var loginResult proto.LoginCmdResult
	loginResult.Code = 200

	userMap := clientMgr.GetAllUsers()

	for userId, _ := range userMap {
		loginResult.User = append(loginResult.User, userId)
	}

	if err != nil {
		loginResult.Code = 500
		loginResult.Error = fmt.Sprintf("%v", err)
	}

	data, err := json.Marshal(loginResult)
	if err != nil {
		fmt.Println("user login result marshal failed", err)
		return
	}

	resMsg.Data = string(data)
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("user login  resMsg result marshal failed", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send message failed", err)
		return
	}
	return

}

func (p *Client) login(msg proto.Message) (err error) {
	defer func() {
		p.loginResp(err)
	}()

	fmt.Printf("recv user.... login request data:%v\n", msg)

	var cmd proto.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}

	fmt.Println("must check user info.")

	_, err = mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}
	clientMgr.AddClient(cmd.Id, p)

	p.NotifyOtherUserOnline(cmd.Id)

	return

}

func (p *Client) NotifyOtherUserOnline(userId int) {
	// 获取在线用户列表
	users := clientMgr.GetAllUsers()

	for id, client := range users {
		if id == userId {
			continue
		}
		fmt.Println("notify user")
		client.NotifyUserOnline(userId)
	}

}

func (p *Client) NotifyUserOnline(userId int) {
	var resMsg proto.Message
	resMsg.Cmd = proto.UserStatusNotifyRes

	var notifyRes proto.UserStatusNotify
	notifyRes.UserId = userId
	notifyRes.Status = proto.UserOnline

	data, err := json.Marshal(notifyRes)
	if err != nil {
		fmt.Println("user notifyRes result marshal failed", err)
		return
	}

	resMsg.Data = string(data)
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("user login resMes result marshal failed", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send message failed", err)
		return
	}

}

func (p *Client) register(msg proto.Message) (err error) {
	var cmd proto.RegisterCmd

	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}

	fmt.Println("start register user..., data:", cmd)

	err = mgr.Register(&cmd.User)
	if err != nil {
		fmt.Println("register failed , err:", err)
		return
	}
	return

}
