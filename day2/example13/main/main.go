package main


import (
	"fmt"
)

func main(){
	var a  int16 = 10
	var b int32

	b = int32(a)
	fmt.Printf("a=%d, b=%d\n", a,b)
}