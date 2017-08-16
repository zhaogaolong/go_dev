package main

import "fmt"
import "time"

func strInput(c chan string, job int) {
	str := "test"
	c <- str
	fmt.Printf("%d-strInput: %s\n", job, str)

}

func strGet(c chan string, job int) {
	str := <-c
	fmt.Printf("%d-Getdata:%s\n", job, str)

}
func main() {
	StrChan := make(chan string, 100)

	for i := 0; i < 100000; i++ {
		// fmt.Println("count-------------------->", i)
		go strInput(StrChan, i)
		go strGet(StrChan, i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("----------Ending-------------")
}
