package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string
	age  int
}

func test(b interface{}) {
	c := reflect.TypeOf(b)
	fmt.Println(c)
	v := reflect.ValueOf(b)
	k := v.Kind()

	fmt.Println(k, v)

}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	c := val.Int()
	fmt.Printf("get value interface{} %d\n", c)
}

func main() {
	a := student{
		name: "aaa",
		age:  121,
	}
	test(a)

	testInt(2324)
}
