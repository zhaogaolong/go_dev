package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Persion struct {
	Name     string
	Age      int
	Describe string
}

var myTemplate *template.Template

func userInfo(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("user info...")
	user := make(map[string]interface{})
	user["name"] = "Conny"
	user["age"] = 18
	user["title"] = "个人网站"

	friendA := Persion{Name: "friendA", Age: 18, Describe: "friendA web"}
	friendB := Persion{Name: "friendB", Age: 18, Describe: "friendB web"}
	friendC := Persion{Name: "friendC", Age: 18, Describe: "friendC web"}
	var friend []Persion

	// 定义friend切片 append 所有的结构体到切片中
	friend = append(friend, friendA, friendB, friendC)

	// 设置user 的friend key 为friend的切片
	user["friend"] = friend

	myTemplate.Execute(writer, user)

}

func initTemplate(fileName string) (err error) {
	myTemplate, err = template.ParseFiles(fileName)
	if err == nil {
		err = fmt.Errorf("parse file err:%v", err)
		return
	}

	return
}

func main() {
	initTemplate("H:/360_update/oldboy_go/src/go_dev/day10/exercises/6-http-template/index.html")
	http.HandleFunc("/user/info", userInfo)

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("server failed ...")
		return
	}

}
