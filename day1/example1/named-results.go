package main

import (
	"fmt"
)

func split(sum int) (x, y int){
	x = sum *4 /9
	y = sum - x
	fmt.Printf("x: %d, y:%d \n", x, y)
	return
}

func main(){
	fmt.Println(split(17))
}