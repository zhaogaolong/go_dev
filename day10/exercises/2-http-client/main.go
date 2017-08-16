package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:8000/user/info")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read data err:", err)
		return
	}
	fmt.Printf("data: %s", string(data))
	// fmt.Printf("data: %v\n", data)
}
