package main

import (
	"fmt"
	"go_dev/day9/exercises/chat/model"
	"go_dev/day9/exercises/chat/proto"
)

var onlineUserMap map[int]*model.User = make(map[int]*model.User)

func outputUserOnline() {
	fmt.Printf("-------------------->userStatus %v\n", onlineUserMap)
	fmt.Println("Online User List:")
	for id, _ := range onlineUserMap {
		if id == userId {
			continue
		}
		fmt.Println("user:", id)
	}

}

func updateUserStatus(userStatus proto.UserStatusNotify) {
	fmt.Printf("------userStatus %v\n", userStatus)
	fmt.Printf("-------------------->onlineUserMap:%v\n", onlineUserMap)
	user, ok := onlineUserMap[userStatus.UserId]
	if !ok {
		user = &model.User{}
		user.UserID = userStatus.UserId
	}
	user.Status = userStatus.Status
	onlineUserMap[user.UserID] = user

}
