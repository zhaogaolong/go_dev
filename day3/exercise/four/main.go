package main

import (
	"fmt"
	"strings"
)

func replace(str, old, new string, count int) string {
	result := strings.Replace(str, old, new, count)
	return result
}

func count(str, word string) int {
	result := strings.Count(str, word)
	return result
}

func repeat(str string, count int) string {
	result := strings.Repeat(str, count)
	return result
}

func main() {
	str := "Hello World! "
	str2 := replace(str, "World", "Conny", 1)
	fmt.Println(str2)

	num := count(str2, "n")
	fmt.Println(num)

	str3 := repeat(str2, 3)
	fmt.Println(str3)

	fmt.Println(strings.ToLower(str3))
	fmt.Println(strings.ToUpper(str3))

}
