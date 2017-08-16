package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputReader *bufio.Reader

func singular(slic []string) {
	/* 这是单数长的字符串 */
	status := true
	long := len(slic)
	middleNumber := len(slic)/2 + 1

	for i := 0; i < long; i++ {
		if i == middleNumber {
			break
		}
		// 如果判断字符串不一致就就标志为false
		if !(slic[i] == slic[(long-i-1)]) {
			status = false
		}
	}
	if status {
		fmt.Println("单数：输入的是回文")
	} else {
		fmt.Println("单数：输入的不是一句回文")
	}

}

func plural(slic []string) {
	status := true
	long := len(slic)
	for i := 0; i < long; i++ {
		// 当运行到中间的数后就不需要对比了
		if i == long%2 {
			break
		}
		// 如果判断字符串不一致就就标志为false
		if !(slic[i] == slic[(long-i-1)]) {
			status = false
			break
		}
	}
	if status {
		fmt.Println("双数：输入的是回文")
	} else {
		fmt.Println("双数：输入的不是一句回文")
	}

}

func ispLalindrome(str string) {
	slic := strings.Fields(str)

	// 设置标志，判断输入的是单数还是复数
	var strLongType bool

	if len(slic)%2 == 0 {
		strLongType = true
	} else {
		strLongType = false
	}
	if strLongType {
		plural(slic)
	} else {
		singular(slic)
	}

}

func main() {
	// var str string
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, _ := inputReader.ReadString('\n')
	fmt.Println("\n")
	fmt.Println(input)
	ispLalindrome(input)
	// ispLalindrome(str)
}
