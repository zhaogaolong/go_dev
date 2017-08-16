package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	n int
}

var m = make(map[int]uint64)
var lock sync.Mutex

func calt(t *task) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		t := &task{n: i}
		go calt(t)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()
	for k, v := range m {
		fmt.Printf("k:%d, v:%d\n", k, v)
	}
	lock.Unlock()

}
