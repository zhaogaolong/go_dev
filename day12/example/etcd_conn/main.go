package main

import (
	"context"
	"fmt"
	"time"

	client "github.com/coreos/etcd/clientv3"
)

func main() {
	cli, err := client.New(client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2379", "localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect faile to etcd")
		return
	}

	fmt.Println("connect etcd success")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	_, err = cli.Put(ctx, "/logagent/conf", "sample_value")
	cancel()
	if err != nil {
		fmt.Println("put failed err:", err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	resp, err := cli.Get(ctx, "/logagent/conf")
	cancel()

	if err != nil {
		fmt.Println("get fail err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Println("ev:", ev.String())
	}

}
