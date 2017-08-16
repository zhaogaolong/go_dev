package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 冒泡排序
func bubbling(slice [20]int) [20]int {
	for i := 0; i < len(slice); i++ {

		// 寻找本次循环的最大值
		for j := 0; j < len(slice)-i-1; j++ {

			if slice[j] > slice[j+1] {
				var1 := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = var1
			}
		}

		fmt.Println(slice)
	}
	return slice

}

func main() {
	var buArray [20]int

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 20; i++ {
		buArray[i] = rand.Intn(500)
	}
	fmt.Println(buArray)
	fmt.Println(bubbling(buArray))

}
