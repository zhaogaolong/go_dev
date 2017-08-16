package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "H:/360_update/oldboy_go/src/go_dev/day7/example/example4/inputFile.txt"
	oututFile := "outputFile.txt"
	buf, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error %s\n", err)
		return
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(oututFile, buf, 0x64)

	if err != nil {
		panic(err.Error())
	}



}
