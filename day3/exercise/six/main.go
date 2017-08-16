package main

import (
	"fmt"
	"strconv"
)

func itoa(num int) (result string) {
	result = strconv.Itoa(num)
	return
}

func atoi(str string) (num int) {
	num, err := strconv.Atoi(str)
	if err == nil {
		return num
	}
	fmt.Printf("Can't convert %s", str)
	fmt.Println(err, "\n")
	return 0
}

func main() {
	num := "w"

	//num2 := itoa(num)
	// fmt.Printf("%s, %T\n", num2, num2)

	num3 := atoi(num)
	fmt.Printf("%d %T\n", num3, num3)
}
