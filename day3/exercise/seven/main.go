package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Printf("%02d/%02d/%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Println(now.Format("02-1-2006 15:04"))
	fmt.Println(now.Format("2006-1-02 15:04"))
	fmt.Println(now.Format("2006-1-02"))
}
