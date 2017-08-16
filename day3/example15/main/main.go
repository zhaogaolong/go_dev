package main

import "fmt"

func main() {
LABEL:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL
			}
			fmt.Printf("i is:%d, and j is:%d\n", i, j)
		}
	}
}
