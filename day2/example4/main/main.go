package main

import (
	"fmt"
	"time"
	
)

const(
	Man = 1
	Female = 2
)

func main(){
	for {
		var s = time.Now().Unix()
		if (s % Female == 0){
			fmt.Println("Female")

		} else{
			fmt.Println("Man")
		}
		time.Sleep(time.Second)

	}
}