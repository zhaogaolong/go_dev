package main

import (
	"fmt"
	"time"
)

func main() {
	startT := time.Now()

	fmt.Println(startT.Format("02/1/2006 15:04"))
	fmt.Println(startT.Format("2006-1-02 15:04"))
	fmt.Println(startT.Format("2006/1/02"))

	endT := time.Now()
	useT := (endT.Nanosecond() - startT.Nanosecond()) / 1000 / 1000
	fmt.Println("use :", useT, "微秒")
}
