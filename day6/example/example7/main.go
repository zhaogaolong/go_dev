package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1 string
	s2 string
	s3 string
}

func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)

	kin := value.Kind()
	fmt.Println(kin)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Fild %d %v\n", i, value.Field(i))
	}

	result := value.Method(0).Call(nil)
	fmt.Println(result)
}
