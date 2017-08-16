package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dsName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("Cant open file: %s\n", srcName)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dsName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Can't create file")
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	CopyFile("file.txt", "outputFile.txt")

}
