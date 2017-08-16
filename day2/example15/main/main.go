package main

import (
	"fmt"
)


func main(){
	var str string = "Hello World!\n"
	var str2 string = `Hello World \n\n\n\n\n\n\`
	fmt.Println(str, str2)
}