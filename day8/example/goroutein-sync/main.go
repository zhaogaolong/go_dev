package main

import "fmt"

func calc(taskChan chan int, resultChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true

		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resultChan <- v
		}
	}
	fmt.Println("exited")
	exitChan <- true
}

func putData(taskChan chan int) {
	for i := 0; i < 10000; i++ {
		taskChan <- i
	}
	close(taskChan)
}

func waitJobsExit(exitChan chan bool, resultChan chan int, jobs int) {

	for i := 0; i < jobs; i++ {
		<-exitChan
	}
	fmt.Println("all jobs exited")
	close(resultChan)
}

func main() {
	jobs := 8
	taskChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go putData(taskChan)

	for i := 0; i < jobs; i++ {
		go calc(taskChan, resultChan, exitChan)

	}
	go waitJobsExit(exitChan, resultChan, jobs)

	for _ = range resultChan {
		//fmt.Println(i)
	}

}
