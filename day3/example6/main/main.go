package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04"))
	fmt.Println(now.Format("2006-1-02 15:04"))
	fmt.Println(now.Format("2006/1/02"))

	end_time := time.Now()

	use_time := (end_time.Nanosecond() - now.Nanosecond()) / 1000
	fmt.Println("use time:", use_time, "微秒")
}
