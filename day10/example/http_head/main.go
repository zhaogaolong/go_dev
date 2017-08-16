package main

import (
	"fmt"
	"net/http"
)

var url = []string{
	"http://www.baidu.com",
	"http://taobao.com",
	"http://google.com",
}

func main() {
	for _, v := range url {

		c := http.Client{
			
		}

		res, err := http.Head(v)
		if err != nil {
			fmt.Println("get head failed, err:", err)
			continue
		}
		fmt.Printf("head success , status:%v\n", res.Status)
	}

}
