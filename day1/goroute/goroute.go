package main

import (
	"fmt"
	"time"
)

func start(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 100; i++ {
		go start(i)
	}
	time.Sleep(time.Second)
}


