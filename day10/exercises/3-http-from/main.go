package main

import (
	"fmt"
	"io"
	"net/http"
)

const from = `
<html>
    <body>
        <form action="#" method="POST" name="bar">
            <input type="text" name="username" />
            <input type="text" name="password" />
            <input type="submit" value="Submit" />
        </form>
    </body>
</html>
`

func simpleServer(w http.ResponseWriter, p *http.Request) {
	io.WriteString(w, "<h1>hello world</h1>")
}

func fromServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch request.Method {
	case "GET":
		io.WriteString(w, from)
	case "POST":
		// 解析form
		request.ParseForm()

		// request.From[username] 返回一个数组，必须要写具体的那个下标的元素
		io.WriteString(w, request.Form["username"][0])
		io.WriteString(w, "\n")

		// request.FromValue 获取一个指定的 name="password"的标签的值
		io.WriteString(w, request.FormValue("password"))
	}
}

func main() {
	http.HandleFunc("/test", simpleServer)
	http.HandleFunc("/from", fromServer)

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("http server failed...")
		return
	}
}
