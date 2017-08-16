package main

import "fmt"

type student struct {
	name string
	six  string
}

func Test(a interface{}) {
	b, ok := a.(student)

	if ok == false {
		fmt.Println("convert failed")
	}

	fmt.Println(b)
}

func just(items ...interface{}) {
	for index, val := range items {
		switch val.(type) {
		case bool:
			fmt.Printf("index %d is bool, value is %v\n", index, val)
		case int, int32, int64:
			fmt.Printf("index %d is int, value is %v\n", index, val)
		case float32, float64:
			fmt.Printf("index %d is float, value is %v\n", index, val)
		case string:
			fmt.Printf("index %d is string, value is %v\n", index, val)
		case student:
			fmt.Printf("index %d is student, value is %v\n", index, val)
		case *student:
			fmt.Printf("index %d is *student, value is %v\n", index, val)
		}

	}

}

func main() {
	var b student
	b.name = "aaa"
	b.six = "666"
	Test(b)

	just(28, 2.2, "ssssss", b, &b, false)
}
