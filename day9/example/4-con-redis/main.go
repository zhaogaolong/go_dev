package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "192.168.11.128:6379")

	if err != nil {
		fmt.Println("Cont redis server failed")
		return
	}
	defer conn.Close()

	// _, err = conn.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println("Set value redis server failed")
		return
	}
	r, err := redis.String(conn.Do("lpop", "book_list"))

	if err != nil {
		fmt.Println("Get value redis server failed")
		return
	}
	fmt.Println(r)

}
