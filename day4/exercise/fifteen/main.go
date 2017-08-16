// 非递归方式实现斐波那契序列

package main

import "fmt"

func fab() {
	var num [100]int

	for i := 0; i <= 10; i++ {
		if i <= 1 {
			num[i] = 1
			fmt.Println(num[i])
		} else {
			num[i] = num[(i-1)] + num[(i-2)]
			fmt.Println(num[i])

		}
	}
}

func main() {
	fab()
}
