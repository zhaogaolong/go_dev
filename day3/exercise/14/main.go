package main

import "fmt"

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operatpr(op add_func, a int, b int) int {
	return op(a, b)
}

func main() {
	c := add
	fmt.Println(c)
	sum := operatpr(c, 2, 5)
	fmt.Println(sum)
}
