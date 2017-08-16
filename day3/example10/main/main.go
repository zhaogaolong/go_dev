package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		str string
	)

	fmt.Scanf("%s", &str)
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println("success convert to int")
	} else {
		fmt.Println("Can't convert to int", num)
	}

}
