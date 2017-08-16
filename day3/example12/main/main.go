package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	rNum := r.Intn(10)
	for {
		fmt.Scanf("%d\n", &num)
		status := false
		switch {
		case rNum == num:
			fmt.Println("success")
			status = true
		case rNum > num:
			fmt.Printf("输入小了\n")
		default:
			fmt.Printf("输入大了\n")
		}
		if status {
			break
		}

	}

}
