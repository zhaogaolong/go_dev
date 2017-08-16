package main

import (
	"fmt"
)

func main() {
	var line int
	fmt.Scanf("%d\n", &line)

	str := "A"
	for i := 1; i <= line; i++ {
		fmt.Println(str)
		str += "A"
	}

}
