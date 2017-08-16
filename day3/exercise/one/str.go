package main

import (
	"fmt"
	"strings"
)

func strHasPrefix(str, head string) string {
	result := strings.HasPrefix(str, head)
	if !result {
		return head + str
	}
	return str
}

func addURL(url string) string {
	if strings.HasSuffix(url, "/") {
		return url
	}

	return url + "/"
}

func firstIndex(str, findWord string) int {
	index := strings.Index(str, findWord)
	return index + 1
}

func lastIndex(str, findWord string) int {
	index := strings.LastIndex(str, findWord)
	return index + 1
}

func main() {
	var (
		url      string
		protocol string
		word     string
	)

	fmt.Scanf("%s%s%s\n", &url, &protocol, &word)
	result := strHasPrefix(url, protocol)
	fmt.Println(result)

	result = addURL(result)
	fmt.Println(result)

	fmt.Printf("word %s in string %s first index %d\n", word, result, firstIndex(result, word))
	fmt.Printf("word %s in string %s last index %d\n", word, result, lastIndex(result, word))

	str1 := "A"
	str1 += "B"
	str1 += "C"
	fmt.Println(str1)

}
