package main

import (
	"fmt"
	"strings"
)

func sTrimSpace(str string) string {
	result := strings.TrimSpace(str)
	return result

}

func sTirm(str, s string) string {
	result := strings.Trim(str, s)
	return result
}

func main() {
	str := "    Hello Conny     "
	str_tp := sTrimSpace(str)
	fmt.Printf("str:%s; trim spaceed:%s\n", str, str_tp)

	str_h := sTirm("He He He He", "He")
	fmt.Printf("%s deled first and last %s result:%s\n ", "He He He He", "He", str_h)

	test_s := "WORLD he he he"
	fmt.Printf("%s deled first word %s result:%s\n ", test_s, "WORLD", strings.TrimLeft(test_s, "WORLD"))
	fmt.Printf("%s deled first word %s result:%s\n ", test_s, "WORLD", strings.TrimRight(test_s, "he"))

	test_s2 := "He,he,he,he "
	fmt.Printf("%s fmt %s result:%s\n ", test_s, "WORLD", strings.Fields(test_s))

	str_slice := strings.Split(test_s2, ",")
	for i := 0; i < len(str_slice); i++ {
		fmt.Println(str_slice[i])
	}
	fmt.Println(strings.Split(test_s2, ","))

}
