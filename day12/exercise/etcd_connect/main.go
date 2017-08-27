package main

import(
	client "github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
)

func main(){
	cli, err := client.New(client.Config{
		Endpoints: []string{"192.168.1.111:2379","192.168.1.111:2379", "192.168.1.111:2379" },
		DialTimeout: 5 * time.Second,
	})
	if err != nil{
		fmt.Println("connect failed , err:", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()

}

