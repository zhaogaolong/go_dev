package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	score float32
}

func main() {
	var stu Student
	stu.Name = "conny"
	stu.Age = 18
	stu.score = 100

	var stu1 *Student = &Student{
		Age:  18,
		Name: "Conny",
	}

	fmt.Printf("%p\n", &stu1.Name)
	fmt.Printf("%p\n", &stu1.Age)
	fmt.Printf("%p\n", &stu1.score)
}
