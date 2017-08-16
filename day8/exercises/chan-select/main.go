package main

import (
	"fmt"
	"time"
)

func test() {
	var ch chan int
	ch = make(chan int, 10)

	ch2 := make(chan int, 10)

	go func() {
		i := 0
		for {
			i++
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i * i
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println("ch", v)
		case v := <-ch2:
			fmt.Println("ch2", v)
		case <-time.After(time.Second):
			fmt.Println("get data timeout")
			// default:
			// 	fmt.Println("get time out ")
			// 	time.Sleep(time.Second)

		}

	}
}

func main() {
	test()

}
