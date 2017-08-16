package main

import (
	"io"
	"os"
)

func copyFile(dsName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dsName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)

}

func main() {
	copyFile("copyFile.txt", "output.txt")

}
