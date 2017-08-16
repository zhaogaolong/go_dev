package main

import (
	"fmt"
	"strings"
)

func stri(str, s string) int {
	result := strings.Index(str, s)
	return result
}

func strend(str, s string) int {
	result := strings.LastIndex(str, s)
	return result
}

func main() {
	str := "hheheheworld"
	first := stri(str, "he")
	last := strend(str, "d")
	fmt.Printf("he在%s中出现的最前的位置是%d\n", str, first)
	fmt.Printf("he在%s中出现的最前的位置是%d\n", str, last)

}
