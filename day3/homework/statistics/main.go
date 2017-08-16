package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inputReader *bufio.Reader

func check(str string) {
	var (
		word   int
		blank  int
		number int
		other  int
	)

	for _, v := range str {
		// fmt.Printf(" %q, %d, index[%d]\n", v, v, i)
		ascii := fmt.Sprintf("%d", v)
		nu, _ := strconv.Atoi(ascii)

		switch {
		case (nu >= 65 && nu <= 90) || (nu >= 97 && nu <= 122):
			// fmt.Printf("word is %q ascii:%d\n", v, nu)
			word++
		case nu == 32:
			blank++
		case nu >= 48 && nu <= 57:
			number++
		default:
			other++
		}

	}
	other -= 2
	fmt.Printf("共有中英文字母：%d\n", word)
	fmt.Printf("共有空格：%d\n", blank)
	fmt.Printf("共有数字：%d\n", number)
	fmt.Printf("共有其他字符：%d\n", other)

}

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, _ := inputReader.ReadString('\n')
	check(input)

}
