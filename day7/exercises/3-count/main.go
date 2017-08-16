package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type count struct {
	letter int
	number int
	space  int
	other  int
}

func test(str string, strCount *count) {
	str1 := []rune(str)
	for _, v := range str1 {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			strCount.letter++
		case v >= '0' && v <= '9':
			strCount.number++
		case v == ' ' || v == '\t':
			strCount.space++
		default:
			strCount.other++
		}
	}
}

func main() {

	file, err := os.Open("os-std.txt")
	if err != nil {
		return
	}
	defer file.Close()
	var strCount count

	reader := bufio.NewReader(file)
	for {

		str, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("read file faild , err:%v", err)
			break
		}
		test(str, &strCount)
	}

	fmt.Println(strCount)

}
