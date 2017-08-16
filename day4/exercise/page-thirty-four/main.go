package main

import "fmt"

func main() {
	var a map[string]string = map[string]string{"hello": "World"}

	a = make(map[string]string, 10)
	a["hello"] = "World"

	val, ok := a["hello"]
	fmt.Println(val, ok)

	for k, v := range a {
		fmt.Println(k, v)

	}
	delete(a, "hello")
	fmt.Println(len(a))
}
