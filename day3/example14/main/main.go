package main

import "fmt"

func main() {
	str := "hello world 中国"

	for i, v := range str {
		if i > 2 {
			continue
		}
		if i > 3 {
			break
		}

		fmt.Printf("index[%d] val[%c] len(%d)\n", i, v, len([]byte(string(v))))
	}

}
