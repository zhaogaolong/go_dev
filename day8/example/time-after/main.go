package main

import (
	"fmt"
	"time"
)

func test() {
	for {
		select {
		case <-time.After(time.Nanosecond):
			fmt.Println("ddd")
		}
	}
}
func main() {

	for i := 0; i < 1024; i++ {

		test()
	}
	time.Sleep(time.Second * 10000)
}
