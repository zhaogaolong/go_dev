package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	reader
	if err != nil {
		fmt.Println("read string faild", err)
	}
	fmt.Printf("input string:%s", str)
}
