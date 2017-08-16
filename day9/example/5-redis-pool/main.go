package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.11.128:6379")
		},
	}

}

func set(str string, val int, exitChan chan bool) {
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("Set", str, val)
	if err != nil {
		fmt.Println("Set error")
		return
	}

	r, err := redis.Int(c.Do("Get", val))
	if err != nil {
		fmt.Println("get efg value failed ")
		return
	}
	fmt.Println(r)
	exitChan <- true

}

func main() {
	exitChan := make(chan bool, 100)

	count := 100

	key := "a"

	for i := 0; i < count; i++ {
		go set(key, i, exitChan)
		key += "a"
	}
	for i := 0; i < count; i++ {
		<-exitChan
	}
	fmt.Println("finsh...")

}
