package main

import "fmt"

func isDaffodil(num int) bool {

	// 求模数
	hundred := num / 100

	//求余数
	result1 := num % 100
	var ten int
	var one int

	if result1 >= 10 {
		ten = result1 / 10
		one = result1 % 10
	} else {
		ten = 0
		one = result1
	}
	// fmt.Println(hundred^3, ten, one)
	if hundred*hundred*hundred+ten*ten*ten+one*one*one == num {
		return true
	} else {
		return false
	}

}

func isNumber(n int) bool {
	var i, j, k int
	i = n % 10

	j = (n / 10) % 10

	sum := i*i*i + j*j*j + k*k*k
	return sum == n
}

func main() {
	for i := 100; i < 1000; i++ {
		if isNumber(i) {
			fmt.Println("水仙花数：", i)
		}
	}
}
