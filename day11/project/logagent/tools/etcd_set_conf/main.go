package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	client "github.com/coreos/etcd/clientv3"
)

type LogConf struct {
	LogPath string `json:"log_path"`
	Topic   string `json:"topic"`
}

var (
	EtcdKey = "/ops/backend/logagent/conf/192.168.11.1"
)

func main() {
	cli, err := client.New(client.Config{
		Endpoints:   []string{"192.168.11.128:2379", "192.168.11.128:2379", "192.168.11.128:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd failed, err:", err)
		return
	}
	defer cli.Close()

	var logConfArr []LogConf
	/*
		logConfArr = append(
			logConfArr,
			LogConf{
				LogPath: "h:/logs/access.log",
				Topic:   "nginx_log",
			},
		)

		logConfArr = append(
			logConfArr,
			LogConf{
				LogPath: "h:/logs/logagent.log",
				Topic:   "nginx_log",
			},
		)
	*/
	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("logConfig json marshal err:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Printf("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
