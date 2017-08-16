package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}

	}()
	panic("dddddddddd")
}

func cale() {
	fmt.Println("bbbbbbbb")
}

func main() {
	go test()
	go cale()
	go cale()
	go cale()
	go cale()
	go cale()
	go cale()
	time.Sleep(time.Second * 10)
}
