package main

import(
	"fmt"
	"day1/goroute_test/goroute"
)

func main(){
	pip := make(chan int, 1)
	go goroute.Add(100, 200, pip)
	sum :=<- pip
	fmt.Println("100 + 200 =", sum)

	go goroute.Sub(200, 100, pip)
	sub :=<- pip
	fmt.Println("200 - 100 =", sub)
}