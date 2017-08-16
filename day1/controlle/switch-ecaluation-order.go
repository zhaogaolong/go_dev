package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("When's Saturday?\n")
	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0 :
			fmt.Println("Today.\n")
		case today + 1 :
			fmt.Println("Tomorrow\n")
		case today + 2:
			fmt.Println("In two days\n")
		default:
			fmt.Print("Too far away. \n")
	}
	fmt.Println(today)
}