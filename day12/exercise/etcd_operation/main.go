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
		fmt.Println("connect etcd failed, err:", err)
		return
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = cli.Put(ctx, "/logagent/conf", "sample_value")
	cancel()
	if err != nil {
		fmt.Printf("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
