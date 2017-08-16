package main

import (
	"fmt"
	ex_a "go_dev/day2/example2/add"
)

func init(){
	ex_a.Test()
	fmt.Println("Initialzed")
}


func main(){
	fmt.Println("Name:", ex_a.Name)
	fmt.Println("Age:",ex_a.Age )
}