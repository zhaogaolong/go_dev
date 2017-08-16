package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error dial", err)
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "q" {
			return
		}
		_, err := conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Println("write sock error")
			return
		}

	}
}
