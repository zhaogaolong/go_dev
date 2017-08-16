package main

import (
	"go_dev/day9/exercises/chat/chat_server/model"
	"time"
)

func main() {
	model.InitRedis("192.168.11.128:6379", 16, 1024, time.Second*300)
	model.InitUserMgr()
	runServer("0.0.0.0:10000")
}
