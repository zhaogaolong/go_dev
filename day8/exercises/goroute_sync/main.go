package main

import (
	"fmt"
)

func calc(taskChan, resChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resChan <- v
		}
	}
	exitChan <- true
}

func read(resChan chan int) {
	for v := range resChan {
		fmt.Println("素数：", v)
	}
}

func putData(taskChan chan int) {
	for i := 0; i < 10000; i++ {
		taskChan <- i
	}
	fmt.Println("put data finsh.")
	close(taskChan)
}

func waitJobs(exitChan chan bool, resChan chan int) {
	for i := 0; i < 8; i++ {
		<-exitChan
	}
	close(resChan)

}

func main() {
	taskChan := make(chan int, 1000)
	resChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go putData(taskChan)
	for i := 0; i < 8; i++ {
		go calc(taskChan, resChan, exitChan)
	}
	go waitJobs(exitChan, resChan)
	read(resChan)

}
