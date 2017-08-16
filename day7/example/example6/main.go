package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outfile, err = os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 066)
	if err != nil {
		fmt.Println("read failed !!")
		return
	}
	defer outfile.Close()

	outfile := bufio.NewReader
}
