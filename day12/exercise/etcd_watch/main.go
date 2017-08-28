package main

import (
	"fmt"
	"time"
	"context"
	client "github.com/coreos/etcd/clientv3"
)

func main(){
	cli, err := client.New(client.Config{
		Endpoints: []string{"192.168.1.111:2379", "192.168.1.111:2379", "192.168.1.111:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil{
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect success.")
	defer cli.Close()

	for {
		rch := cli.Watch(context.Background(), "/logagent/conf")
		for wresp := range rch{
			for _, ev := range wresp.Events{
				fmt.Printf("%s  %q: %q \n", ev.Type, ev.Kv.Key, ev.Kv.Value         )
			}
		}
	}
}