package main

import "fmt"

func modify(arr *[5]int) {
	(*arr)[3] = 100
	return
}

func main() {
	var a [5]int
	modify(&a)
	fmt.Println(a)

}
