package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index hello")
	fmt.Fprintf(w, "index hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user login")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("server failed, err:", err)
	}
}
