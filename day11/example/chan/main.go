package main

import (
	"fmt"
)

func test(int) <-chan int {
	chan1 := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {

			chan1 <- 1
		}
		close(chan1)

	}()
	return chan1

}

func main() {
	chan1 := make(chan string, 10)
	chan2 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		chan1 <- "test"
		chan2 <- i
	}

	for i := 0; i < 20; i++ {
		select {
		case msg := <-chan2:
			fmt.Println("chan2", msg)
		case msg := <-chan1:
			fmt.Println("chan1", msg)
		}
	}
}
