package main

import (
	"fmt"
	"strconv"
)

func convert(str string) {
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("Convert number:%d", num)
	} else {
		fmt.Printf("Can't convert to int, ERROR:%s", err)
	}
}

func main() {
	var (
		sNum string
	)
	fmt.Scanf("%s", &sNum)

	convert(sNum)
}
