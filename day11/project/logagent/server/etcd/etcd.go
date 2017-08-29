package etcd

import (
	"context"
	"fmt"
	"time"

	client "github.com/coreos/etcd/clientv3"
)

var (
	cli     *client.Client
	InitCli bool
)

func InitEtcd(addr string) {
	cli, err := client.New(client.Config{
		Endpoints:   []string{addr},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Println("connect etcd failed, err:", err)
		return
	}
	InitCli = true
	// defer cli.Close()

}

func RunEtcd() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := cli.Put(ctx, "/logagent/conf", "sample_value")
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

func updateConfig() {}

func WatchEtcd() {}
