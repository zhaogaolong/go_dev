package main

import (
	"fmt"
	"io"
	"net/http"
)

const from = `<html> <boby>
<form action="#" method="post">
  <input type="text" name="in" />
  <input type="text" name="in" />
  <input type="submit" value="Submit" />
</form>
</boby>
</html>
`

func simpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "hello world")

}

func fromServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Println("request: method", request.Method)
	switch request.Method {
	case "GET":
		io.WriteString(w, from)
		fmt.Println("get")
	case "POST":
		fmt.Println("posttt")
		request.ParseForm()
		io.WriteString(w, request.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("in"))
	}

}

func logPanincs(handle http.HandleFunc) http.HandleFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				fmt.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}

		}()
		handle(w, request)

	}
}

func main() {
	http.HandleFunc("/test1", simpleServer)
	http.HandleFunc("/test2", fromServer)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)

	}
}
