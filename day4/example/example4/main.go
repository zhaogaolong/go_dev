package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func main() {
	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("test"))
	fmt.Println(f2("pic"))
}
