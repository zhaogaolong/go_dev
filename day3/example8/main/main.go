package main

import "fmt"

func test(i *int) {
	fmt.Println(*i)
	var p *int = 100
	

	*i = 9
	fmt.Println(*i)

}

func main() {
	a := 10
	fmt.Println("a address:", &a)
	fmt.Println("a address:", a)

	test(&a)

}
