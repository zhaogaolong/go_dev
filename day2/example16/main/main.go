package main

import (
	"fmt"
)

func main() {
	var str string = "hello world"
	var str2 string
	for i := 0; i < len(str); i++ {
		str2 = str2 + fmt.Sprintf("%c", str[len(str)-1-i])
	}
	fmt.Println(str2)
}
