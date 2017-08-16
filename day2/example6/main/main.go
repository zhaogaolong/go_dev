package main

import (
	"fmt"
	_ "time"
)

func modify(a *int){
	*a = 10
}


func main(){
	var a = 100
	var b chan int = make(chan int, 1)

	fmt.Println("b:", b)

	fmt.Println("a:", a)
	modify(&a)
	fmt.Println("a:", a)
}