package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeEnd() {
	file, err := os.OpenFile("book.db", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Open file error")
	}
	defer file.Close()

	file.WriteString("test string\n")
	file.WriteString("test string\n")
}

func readLine() {
	file, err := os.Open("book.db")
	if err != nil {
		fmt.Println("Open file error")
	}
	defer file.Close()
	buff := bufio.NewReader(file)
	for {
		str, err := buff.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("str:-->%s", str)

	}

}

func main() {

	readLine()
}
