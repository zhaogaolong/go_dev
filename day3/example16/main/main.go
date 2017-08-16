package main

import "fmt"

type add_func 
func add(a, b int) int {
	return a + b
}

func main() {
	c := add
fmt.Println(c)
	sum := c(3, 5)
	fmt.Println(sum)

}
