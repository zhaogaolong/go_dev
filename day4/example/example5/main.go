package main

import "fmt"

func main() {
	var a [10]int
	a[2] = 100
	a[4] = 400

	b := a

	b[2] = 444

	fmt.Println(a)
	fmt.Println(b)
}
