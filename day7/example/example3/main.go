package main

import (
	"bufio"
	"fmt"
	"os"
)

type count struct {
	word   int
	number int
	other  int
	none   int
}

func (p *count) test(str string) {

	// runArray := []rune(str)

	for i := 0; i < len(str); i++ {
		s := str[i]
		switch {
		case s >= 'a' && s <= 'z':
			fallthrough
		case s >= 'A' && s <= 'Z':
			p.word++
		case s == ' ':
			p.none++
		case s >= '0' && s <= '9':
			p.number++
		default:
			p.other++
		}
	}

}

func main() {
	file, _ := os.Open("H:/360_update/oldboy_go/src/go_dev/day7/example/example3/test")
	read := bufio.NewReader(file)
	// str, _ := read.ReadString('\n')
	str, _, _ := read.ReadLine()

	// var c count
	// c.test(str)
	fmt.Println(str)
	fmt.Printf("%T", str)
	file.Close()
}
