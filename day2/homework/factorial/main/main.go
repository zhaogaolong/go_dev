package main

import (
	"fmt"
)

func factorial(number int) uint64 {
	var result uint64 = 1
	var i uint64 = 2
	var count = uint64(number)
	for ; i <= count; i++ {
		result = result * i

	}
	fmt.Printf("%d 的阶乘是：%d\n", number, result)
	return result
}

func main() {
	factorial(20)
}
