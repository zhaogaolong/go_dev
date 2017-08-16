package main

import "fmt"

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	fmt.Println("nnn--->", n)
	return fab(n-1) + fab(n-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("iiiiii------------->", i)
		n := fab(i)
		fmt.Println(n)
	}
}
