package main

import (
	"fmt"
	"os"
	"text/template"
)

type Persion struct {
	Name string
	Age  int
}

func main() {

	file := "h:/360_update/oldboy_go/src/go_dev/day10/exercises/http_template/index.html"
	t, err := template.ParseFiles(file)

	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Persion{
		Name: "Conny",
		Age:  18,
	}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
