package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename := "my.log"

	tailfs, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,  // tail -F 重新打开
		Follow:    true,  // 更新后继续读
		MustExist: false, // 必须存在，true：如果不存在就退出，false：如果没有就等待检查。
		// Location: &tail.SeekInfo{Offset:0, Whence:2},
	})

	if err != nil {
		fmt.Println("tailf file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tailfs.Lines
		if !ok {
			fmt.Println("tailf chan closed reopen, filename:", filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
