package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Print("Go runs on \n")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "Linux":
			fmt.Println("Linux")
		case "windows":
			fmt.Println("Run on Windows system.")
		default:
			fmt.Printf("%s.", os)
	}
}