package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() {
	res, err := http.Get("http://localhost:8000")

	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read failed, err:", err)
		return
	}
	fmt.Println("data")
	fmt.Println(data)

}

func main() {
	go get()

}
