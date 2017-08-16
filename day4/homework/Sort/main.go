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

// 选择排序（selection sort）从小到大的
func selection(slice [20]int) [20]int {
	for i := 0; i < len(slice); i++ {
		index := i

		for j := i; j < len(slice); j++ {
			if index == len(slice)-1 {
				break
			}

			// find least number
			if slice[index] > slice[j] {
				index = j
			}

		}
		fmt.Println("最大数的index:", index, i)
		if slice[i] != slice[index] {
			tmpN := slice[i]
			slice[i] = slice[index]
			slice[index] = tmpN

		}

	}
	return slice

}

func add(a []int, b []int) {
	a = append(a, b...)
}

// 插入排序
func insertNew(slice []int) []int {
	var sortSlice []int
	count := len(slice)
	for i := 0; i < count; i++ {
		tmpN1 := slice[0]
		slice = slice[1:]

		if len(sortSlice) == 0 {
			sortSlice = append(sortSlice, tmpN1)
		} else {
			for j := 0; j < len(sortSlice); j++ {

				// 能找到比自己大的值
				if sortSlice[j] > tmpN1 {
					tmpN2 := sortSlice[j:]
					sortSlice = append(sortSlice[:j], tmpN1)
					sortSlice = append(sortSlice, tmpN2...)
					break
				} else if j == len(sortSlice)-1 {
					// 找不到比自己大的值就放在最后
					sortSlice = append(sortSlice, tmpN1)
					break
				}

			}

		}
		// fmt.Println(i, count, "\t", sortSlice, slice)

	}
	return sortSlice
}

func qsort(a []int, left, right int) {
	if left > right {
		return
	}

	val := a[left]
	k := left

	for i := left + 1; i <= right; i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}

	a[k] = val
	qsort(a, left, k-1)
	qsort(a, k+1, right)
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	/*
		var buArray [20]int
		for i := 0; i < 20; i++ {
			buArray[i] = rand.Intn(500)
		}
		fmt.Println(buArray)

		冒泡排序
		fmt.Println(bubbling(buArray))

		选择排序
		fmt.Println(selection(buArray))
	*/

	var slice []int
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 20; i++ {
		slice = append(slice, rand.Intn(500))
	}
	fmt.Println(slice)
	qsort(slice[:], 0, len(slice)-1)
	fmt.Println(slice)
	// fmt.Println(insertNew(slice))

}
