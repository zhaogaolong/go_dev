package main

import "fmt"

type student struct {
	Name  string `json "student_name"`
	Age   int
	Score float32
	Sex   string
}

func (s student) Print() {
	fmt.Println("----start--------")
	fmt.Println(s)
	fmt.Println("----end--------")

}


func (s student) Set(name string, age int, score float32, sex string){
	s.Name = name
	s.Age = age
	s.Score score
	s.Sex = sex
}
