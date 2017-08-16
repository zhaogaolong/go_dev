package main

import "fmt"

func change(num []int) {
	num[0] = 100

}

func array() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	change(slice)
	fmt.Println(slice[:])
	fmt.Println(slice[4:])
	fmt.Println(slice[4:8])
	fmt.Println(slice[:len(slice)])
}

func app() {
	var a = []int{1, 2, 3}
	var b = []int{4, 9, 6}
	a = append(a, b...)
	fmt.Println(a)

	for index, val := range a {
		fmt.Println(index, val)
	}

}

func copy() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 10)
	fmt.Println(s1)
	fmt.Println(s2)

	// copy(s2, s1)
	fmt.Println(s2)

	s3 := []int{1, 2, 3}
	s3 = append(s3, s2...)
	fmt.Println(s3)
}

func str() {
	str := "hello world"
	s := []byte(str)
	s[0] = 'o'
	str = string(s)
	fmt.Println(str)
}

func main() {

	// array()
	// app()
	// copy()
	str()
}
