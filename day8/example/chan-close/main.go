package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

	for {
		b, ok := <-ch
		if !ok {
			fmt.Println("chan is closed.")
			break
		}
		fmt.Println(b)

	}
}
