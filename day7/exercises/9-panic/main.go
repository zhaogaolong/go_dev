package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\n", e)
		}
	}()
	badCall()
	fmt.Println("After bad call")
}

func main() {
	fmt.Println("Calling test")
	test()
}
