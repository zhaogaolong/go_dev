package proto

import "go_dev/day9/exercises/chat/model"

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

type RegisterCmd struct {
	User model.User `json:"user"`
}

type LoginCmdResult struct {
	Code  int    `json:"code"`
	User  []int  `json:"users"`
	Error string `json:"error"`
}

type UserStatusNotify struct {
	UserId int `json:"user_id"`
	Status int `json:"user_status"`
}
