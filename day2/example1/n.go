package main

import (
	"fmt"
)


func main(){
	print_number(5)

}

func print_number(n int) {
	for x := 0; x<=n; x++ {
		fmt.Printf("%d+%d=%d\n", x, n-x, n)
	}

}