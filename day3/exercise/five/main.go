package main

import (
	"fmt"
	"strings"
)

func trimSpace(str string) (result string) {
	result = strings.TrimSpace(str)
	return result
}

func trimLift(str string) (result string) {
	result = strings.TrimLeft(str, " ")
	return result
}

func trimRight(str string) (result string) {
	result = strings.TrimRight(str, " ")
	return result
}

func split(str, word string) (result []string) {
	result = strings.Split(str, word)
	return result
}

func main() {
	str := "       Hello World    "

	str2 := trimSpace(str)
	fmt.Println(str2, "\n")

	str2 = trimLift(str)
	fmt.Println(str2, "\n")

	str2 = trimRight(str)
	fmt.Println(str2)

	str2 = trimSpace(str)

	result := split(str2, "o")
	fmt.Println(result)

	str2 = strings.Join(result, "o")
	fmt.Println(str2)

}
