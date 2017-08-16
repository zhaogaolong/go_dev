package model

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func InitRedis(addr string, idelConn, maxCon int, idelTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     idelConn,
		MaxActive:   maxCon,
		IdleTimeout: idelTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
	return
}

func GetConn() redis.Conn {
	return pool.Get()
}

func PutCon(conn redis.Conn) {
	conn.Close()
}
