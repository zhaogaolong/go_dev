package main

import "fmt"

func main() {
	var i = 0
	switch i {
	case 0:
		fallthrough
	case 1, 3, 4, 5, 6:
		fmt.Println("1")
	case 7, 8, 9:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}
}
