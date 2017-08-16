package main

import "fmt"

type Car1 struct {
	name string
	age  int
}

type Car2 struct {
	name string
	age  int
}

type Trainb struct {
	Car1
	Car2
}

func main() {
	var t Trainb
	t.Car1.name = "test"
	t.Car1.age = 10

	fmt.Println(t)
}
