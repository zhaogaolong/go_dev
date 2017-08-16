package main

import (
	"fmt"
	"net/http"
)

func test1(w http.ResponseWriter, p *http.Request) {
	fmt.Println("get index")
	panic("server error")
}

func logPanic(handle http.HandlerFunc) http.HandlerFunc {

	return func(write http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				fmt.Printf("[%v] caught panic:%v\n", request.RemoteAddr, x)
			}

		}()
		handle(write, request)
	}
}

func main() {

	// http.HandleFunc("/test", logPanic(test1))
	http.HandleFunc("/test", test1)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("server license err")
	}
}
