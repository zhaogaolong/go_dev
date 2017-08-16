package main

import (
	"fmt"
	"strings"
)

func sReplace(str, old, new string, n int) string {
	str = strings.Replace(str, old, new, n)
	return str
}

func sCount(str, s string) int {
	number := strings.Count(str, s)
	return number
}

func sRepeat(str string, count int) string {
	result := strings.Repeat(str, count)
	return result
}

func main() {
	str := "Hello World, Hello World"
	str = sReplace(str, "World", "Conny", 1)
	fmt.Printf("new str: %s\n", str)

	count := sCount(str, "Hello")
	fmt.Printf("str %s in %s 出现%d次\n", "Hello", str, count)

	str_repeat := sRepeat(str, 2)
	fmt.Printf("%s REPEAT %d:%s\n", str, 2, str_repeat)

	fmt.Printf("Lower: %s\n", strings.ToLower(str))
	fmt.Printf("Upwer: %s\n", strings.ToUpper(str))
}
