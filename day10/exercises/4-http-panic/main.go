package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("page --> index")
	fmt.Fprintln(writer, "<h1>welcome index</h1>")
}

func login(writer http.ResponseWriter, request *http.Request) {
	panic("user logiun page panic...")

}
func logPanic(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				fmt.Printf("[%v] caught panic: %v\n", request.RemoteAddr, x)
			}

		}()
		handle(writer, request)

	}
}

func main() {
	http.HandleFunc("/", logPanic(index))
	http.HandleFunc("/login", logPanic(login))
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("server failed ...")
		return
	}
}
