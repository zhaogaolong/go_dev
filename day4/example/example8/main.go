package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1, 8, 6, 3, 5, 7, 0, 1, 32, 6}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testString() {
	var a = [...]string{"o", "dw", "zeer", "az"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{5.21, 43.1, 23, 3465.5, 234, 745.0}
	sort.Float64s(a[:])
	fmt.Println(a)

}

func testIntSearch() {
	var a = [...]int{1, 8, 6, 3, 5, 7, 0, 1, 32, 6}
	sort.Ints(a[:])
	index := sort.SearchInts(a, 11)
	fmt.Println(index)
}
func main() {
	testIntSort()
	testString()
	testFloat()
}
