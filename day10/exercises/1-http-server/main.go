package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, p *http.Request) {
	fmt.Println("index ....")
	fmt.Fprintln(w, "<h1>index</h1>")
}

func login(w http.ResponseWriter, p *http.Request) {
	fmt.Println("login")
	fmt.Fprintln(w, "login")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("http server failed...")
		return
	}
}
