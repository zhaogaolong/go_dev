package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Scort int
}

func (p *Student) init(name string, age, scort int) {
	p.Age = age
	p.Name = name
	p.Scort = scort
	fmt.Println(p)
}

func (p Student) get() Student {
	return p
}

func main() {
	var stu Student

	stu.init("hh", 18, 80)

	stu1 := stu.get()
	fmt.Println(stu1)
}
