package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println("hello", v)
		t.Stop()
	}

	// select {
	// case <-time.After(time.Second):
	// 	fmt.Println("after")
	// }

}
