package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go_dev/day11/project/logagent/server"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	client "github.com/coreos/etcd/clientv3"
)

var (
	EtcdClient *client.Client
	Collect    []server.CollectConf
)

func InitEtcd(addr, etcdKey string) (err error) {
	for {
		err := connectEtcd(addr)
		if err != nil {
			logs.Error("connect etcd faile, err:", err)
		} else {
			logs.Debug("connect etcd success")
			break
		}
		time.Sleep(time.Second)
	}
	getConfig(etcdKey)
	return
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
	EtcdClient = cli
	return
}

func getConfig(key string) (err error) {

	// 是否以／结尾，如果不是就添加一个
	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}
	for _, ip := range localIPArry {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := EtcdClient.Get(ctx, etcdKey)
		logs.Debug("get etcd config key:%s", etcdKey)
		if err != nil {
			continue
		}
		cancel()
		for _, v := range resp.Kvs {
			var colConf = []server.CollectConf{}

			err := json.Unmarshal(v.Value, &colConf)
			if err != nil {
				logs.Error("etcd config umarshal failed, err: %v", err)
				continue
			}
			logs.Info("etcd config umarshal success, %v", colConf)

			logs.Info("etcd config : %s", v.Value)
		}

	}
	return
}

func RunEtcd() {

}

func updateConfig() {}

func WatchEtcd() {}
