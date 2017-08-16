package main

import (
	"fmt"
)

func fab(a [100]int64) {
	for i := 0; i < len(a); i++ {
		if i == 0 || i == 1 {
			a[i] = 1
		} else {
			a[i] = a[(i-1)] + a[(i-2)]
		}
	}
	fmt.Println(a)

}

func testArray() {
	var a [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	for row, v := range a {
		for col, v1 := range v {
			fmt.Printf("(%d, %d)=%d ", row, col, v1)
		}
		fmt.Println()
	}
}

func main() {
	// var a [100]int64
	// fab(a)
	testArray()
}
