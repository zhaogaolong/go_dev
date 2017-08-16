package main

import (
	"fmt"
	"strings"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
		fmt.Printf("urlProcess %s\n", url)
		return url
	}
	return url

}

func pathProcess(path string) string {
	if !strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s/", path)
	}
	return path

}

func main() {
	var (
		url  string
		path string
	)
	fmt.Scanf("%s%s", &url, &path)

	fmt.Printf("intput values,url:%s, path:%s\n", url, path)

	url = urlProcess(path)
	fmt.Printf("right url: %s\n", url)

	path = pathProcess(url)
	fmt.Printf("right path: %s\n", path)

}
