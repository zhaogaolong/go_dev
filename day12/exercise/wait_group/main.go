package main

import (
	"fmt"
	"sync"
	"time"
)

func calc(w *sync.WaitGroup, i int) {
	w.Add(1)
	fmt.Println("calc", i)
	time.Sleep(time.Second)
	w.Done()
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go calc(&wg, i)
	}
	wg.Wait()
}
