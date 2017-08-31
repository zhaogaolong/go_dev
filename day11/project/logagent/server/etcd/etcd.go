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

	// 主要用于当配置改变后通知
	EtcdConfigUpdate chan int
)

func InitEtcd(addr, etcdKey string) (err error) {
	// 初始化更新管道
	EtcdConfigUpdate = make(chan int, 1)

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
		// logs.Debug("get etcd config key:%s", etcdKey)
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
			Collect = colConf

			// start watch key
			go watchEtcd(etcdKey)

			logs.Info("etcd config umarshal success, %v", colConf)

			logs.Info("etcd config : %s", v.Value)
		}

	}
	return
}

func RunEtcd() {

}

func updateConfig() {}

func watchEtcd(key string) {
	time.Sleep(time.Second * 2)
	for {
		rch := EtcdClient.Watch(context.Background(), key)
		for _ = range rch {
			// fmt.Println("update config ", wresp)
			logs.Warn("[etcd watch] etcd config changed masu update config")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			resp, err := EtcdClient.Get(ctx, key)
			if err != nil {
				logs.Error("[etcd watch]  get new config failed, err: %v", err)
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
				Collect = colConf
			}
			logs.Warn("[etcd watch] update config %v", Collect)
			EtcdConfigUpdate <- 1
		}
	}
}
