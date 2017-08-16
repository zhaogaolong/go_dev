package main

import "fmt"

func perfectNumber(num int) {
	var sum int
	for i := 1; i < num; i++ {
		remainder := num % i
		if remainder == 0 {
			sum += i

		}

	}
	if sum == num {
		fmt.Printf("perfect number:%d\n", num)
	}
}
func main() {
	for i := 1; i < 1000; i++ {
		perfectNumber(i)
	}
	for true {
		fmt.Println("6666666666\n")
	}
}
