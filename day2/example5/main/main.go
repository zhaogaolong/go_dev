package main

import (
	"fmt"
	"os"
	"runtime"
)

func main(){
	var goos = os.Getenv("GOOS")	
	fmt.Printf("SYStem:%s\n", goos)
	var hostname = runtime.Goos()

	fmt.Printf("SYStem:%s\n", hostname)
	var path = os.Getenv("PATH")
	fmt.Printf("PATH is %s\n", path)
}