package main

import (
	"fmt"
	"go_dev/day11/project/logagent/server/etcd"
	"go_dev/day11/project/logagent/server/tailf"

	"github.com/astaxie/beego/logs"
)

func initConfig() (err error) {
	configPath := "H:/360_update/oldboy_go/src/go_dev/day11/project/logagent/conf/logagent.conf"
	err = loadConf("ini", configPath)
	if err != nil {
		fmt.Println("load config failed, err:", err)
		panic("load config failed")
		return
	}
	return
}

func initRun() {
	err := etcd.InitEtcd(appConf.etcdAddr, appConf.etcdKey)
	if err != nil {
		logs.Warn("Init tailf err, err:%v", err)
		return
	}

	err = tailf.InitTailf()
	if err != nil {
		logs.Warn("Init tailf err, err:%v", err)
		return
	}
	// kafka.ServerRun(appConf.kafka_addr)

}

func main() {
	err := initConfig()
	if err != nil {
		fmt.Println("config read failed, err:", err)
		return
	}
	initRun()
	logs.Debug("initialize all success")
}
