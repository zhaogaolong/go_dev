package main


import (
	"fmt"
)

const (
	a = iota
	b
	c
	d
	e
	f = iota
	g
	h
	i
)


func main(){

	fmt.Println(a,b,c,d,e, f,g,h,i)
}