package main

import "fmt"

func add(a int, arg ...int) (sum int) {
	sum = a
	for i := 0; i < len(arg); i++ {
		sum = sum + arg[i]
	}
	return

}

func concat(arg ...string) (str string) {
	for i := 0; i < len(arg); i++ {
		str = str + arg[i]
	}
	return

}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(concat("H", "e", "l", "l", "o"))
}
