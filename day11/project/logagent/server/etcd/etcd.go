package etcd

import (
	"strings"
	"context"
	"fmt"
	"time"
	"go_dev/day11/project/logagent/server"
	"github.com/astaxie/beego/logs"
	client "github.com/coreos/etcd/clientv3"
)

var (
	EtcdClient	*client.Client
	Collect 	[]server.CollectConf
)

func InitEtcd(addr , etcdKey string) {
	for {
		err := connectEtcd(addr)
		if err != nil {
			logs.Error("connect etcd faile, err:", err)
		} else {
			break
		}
		time.Sleep(time.Second)
	}
	getConfig(etcdKey)
	// startTailf()

}
func connectEtcd(addr string) (err error) {
	cli, err := client.New(client.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return
	}
	return
}

func getConfig(key string) (err error) {

	// 是否以／结尾，如果不是就添加一个
	if strings.HasSuffix(key, "/") == false{
		key = key + "/"
	}

	for _, ip := range localIPArry{
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := EtcdClient.Get(ctx, etcdKey)
		if err != nil{
			continue
		}
		cancel()
		for k, v := range resp.Kvs{
			fmt.Println(k, v)
		}

	}

	return
}

func RunEtcd() {

}

func updateConfig() {}

func WatchEtcd() {}
