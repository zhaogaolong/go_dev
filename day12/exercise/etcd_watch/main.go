package main

import (
	"context"
	"fmt"
	"time"

	client "github.com/coreos/etcd/clientv3"
)

func main() {
	cli, err := client.New(client.Config{
		Endpoints:   []string{"192.168.11.128:2379", "192.168.11.128:2379", "192.168.11.128:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect success.")
	defer cli.Close()

	for {
		// 获取一个watch chan
		rch := cli.Watch(context.Background(), "/logagent/conf")

		// 循环管道，获取一个watchResponse
		for wresp := range rch {

			// 循环watchResponse 获取事件，也就是最终的值
			for _, ev := range wresp.Events {
				fmt.Printf("%s  %q: %q \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}
