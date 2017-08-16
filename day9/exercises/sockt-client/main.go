package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("error dialing, err:", err)
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
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Println("sockt write failed, err:", err)
			return
		}

	}

}

func main() {
	client()
}
