package main

import "fmt"

func nineTables() {
	var str string
	for i := 1; i < 10; i++ {
		str = fmt.Sprintf("%dx%d=%d", 1, i, 1*i)

		for j := 2; j <= i; j++ {
			str += fmt.Sprintf("| %dx%d=%d", j, i, j*i)
		}
		fmt.Print(str, "\n")
	}
}

func main() {
	nineTables()
}
