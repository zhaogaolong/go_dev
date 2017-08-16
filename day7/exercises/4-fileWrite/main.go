package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputErr := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if outputErr != nil {
		fmt.Println("An error occured with file crea iond on")
		return
	}

	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
