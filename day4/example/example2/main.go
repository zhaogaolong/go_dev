package main

import "fmt"

func test() {
	s1 := new([]int)
	fmt.Println(s1)

	s2 := make([]int, 10)
	fmt.Println(s2)

	*s1 = make([]int, 5)
	(*s1)[0] = 100
	s2[0] = 100
	fmt.Println(*s1)
	fmt.Println(s2)

	var s3 []string
	var s4 = []string{"1", "2", "3", "4"}

	s3 = append(s3, "test")
	s3 = append(s3, s4...)

	fmt.Println(s3)

	var arrKeyValue1 = [5]string{3: "Chris", 4: "Ron"}
	var arrLazy2 = [...]int{5, 6, 7, 8, 22}
	var arrLazy3 = []int{5, 6, 7, 8, 22}
	var arr10 []int = arrLazy3[:3]

	arr10 = append(arr10)

	fmt.Printf("%T\n", arrKeyValue1)
	fmt.Printf("%T\n", arrLazy2)
	fmt.Printf("%v %T\n", arrLazy3, arrLazy3)

	fmt.Printf("%v %T\n", arr10, arr10)
}

func main() {
	test()
}
