package main

import "fmt"

func sendData(ch chan<- string) {
	ch <- "Beijing"
	ch <- "London"
	ch <- "Tokio"
	ch <- "Taiwan"
	ch <- "other"
	ch <- "Pairs"
	close(ch)
}

func getData(ch <-chan string, exitChan chan bool) {
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("getData exited")
	exitChan <- true
	close(exitChan)
}

func main() {
	jobs := 8
	ch := make(chan string, 10)
	exitChan := make(chan bool, jobs)
	go sendData(ch)
	go getData(ch, exitChan)
	for _ = range exitChan {
	}

}
