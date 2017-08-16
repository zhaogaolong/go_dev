package main

import (
	"fmt"
)

func isPrime(x int) bool {
	if x == 1 {
		return false
	}

	// 循环2到小于自己的数判断
	for i := 2; i < x; i++ {
		if x%i == 0 {
			//不是素数
			return false
		}

	}
	return true

}

func main() {
	// 循环2-100内所有的数字
	for i := 1; i <= 1000; i++ {
		if isPrime(i) {
			fmt.Println("素数：", i)
		}
	}
}
