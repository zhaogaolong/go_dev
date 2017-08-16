package main

import "fmt"

func main() {
	ch := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		ch <- i

	}
	// close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
