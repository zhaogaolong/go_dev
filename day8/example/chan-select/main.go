package main

import (
	"fmt"
	"time"
)

func send(ch1 chan int, ch2 chan int) {
	var i int
	for {
		ch1 <- i
		ch2 <- i * i
		time.Sleep(time.Second)
		i++
	}
}

func main() {
	var ch1 chan int
	var ch2 chan int
	ch1 = make(chan int, 10)
	ch2 = make(chan int, 10)
	go send(ch1, ch2)
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case s := <-ch2:
			fmt.Println(s)
		default:
			fmt.Println("time out ")
			time.Sleep(time.Second)
		}
	}

}
