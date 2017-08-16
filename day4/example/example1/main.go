package main

import (
	"errors"
	"fmt"
	"time"
)

func initConfig() (err error) {
	return errors.New("init config faild")
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	for {
		test()
		fmt.Println("hello")
		time.Sleep(time.Second)
	}

}
