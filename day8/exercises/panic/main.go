package main

import "fmt"
import "time"

func task() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}

	}()

	var m map[string]int

	m["mus"] = 10

}

func main() {
	for i := 0; i < 100; i++ {
		go task()
	}
	time.Sleep(time.Second)
}
