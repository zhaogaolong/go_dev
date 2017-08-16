package main

import "fmt"

type student struct {
	Name string
	Info map[string]int
}

func main() {
	stu1 := student{
		Name: "stu1",
		Info: make(map[string]int),
	}

	stu1.Info["book1"] = 5

	if stu1.Info == nil {
		fmt.Println("stu1 info is nil")
	} else {
		fmt.Println("stu1 info is not nil")
	}

	if _, ok := stu1.Info["book2"]; ok {
		fmt.Println("has book2")
	} else {
		fmt.Println("hasn't book2", ok)
	}

}
