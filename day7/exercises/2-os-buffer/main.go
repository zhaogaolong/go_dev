package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	file, err := os.Open("os-std.txt")
	if err != nil {
		return
	}
	defer file.Close()
	inputReader = bufio.NewReader(file)
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s", input)
	}
}
