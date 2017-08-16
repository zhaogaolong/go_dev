package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Person struct {
	Name  string
	Title string
	Age   int
}

var tem *template.Template

var user *Person

func hello(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("H:/360_update/oldboy_go/src/go_dev/day10/example/http_template/index.html")
	if err != nil {
		fmt.Println("parse failed... ")
		return
	}
	user = &Person{
		Name:  "名字",
		Title: "个人网站",
		Age:   18,
	}

	fmt.Println("index hello")
	tem.Execute(w, user)

}

func main() {

	http.HandleFunc("/user/login", hello)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("server failed, err:", err)
	}

}
