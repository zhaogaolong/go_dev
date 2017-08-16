package main

import (
	"fmt"
)

func revers(x, y int) (int, int){
	return y, x	
}

func main(){
	a := 3
	b := 4
	a, b = revers(a, b)
	fmt.Print(a,b)
}
