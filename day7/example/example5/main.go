package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func ReadZip() {
	fName := "myfile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	defer fi.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v Cant't Opent file %s", os.Args[0], fName)
	}

	fz, err := gzip.NewReader(fi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v Cant't Opent file %s", os.Args[0], fName)
	}

	r = bufio.NewReader(fz)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)

	}

}

func main() {
	ReadZip()

}
