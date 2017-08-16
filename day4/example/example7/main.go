package main

import (
	"fmt"
)

func change(a []int) {
	a[0] = 100
}

func test() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[1:2]

	fmt.Printf("%p\n", &a[1])
	fmt.Printf("%p\n", b)

	change(b)
	fmt.Println(a)
}

func testC() {
	str := "Hello World"
	s := []rune(str)
	s[0] = '0'
	str = string(s)
	fmt.Println(str)
}

func main() {
	testC()
}
