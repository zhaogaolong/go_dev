package main

import (
	"github.com/astaxie/beego/logs"
	"fmt"
)


func main(){

	// 初始化配置
	err := initConfig("ini", "/Users/zhaogaolong/develop/go_dev/src/go_dev/day13/example/log_transfer/conf/log_transfer.conf")
	if err != nil{
		panic(err)
	}

	fmt.Println(appConfig)
	// 初始化log
	err = initLogger(appConfig.LogLevel, appConfig.LogPath)
	if err != nil{
		panic(err)
	}

	// 初始化kafka
	err = initKafka(appConfig.kafkaAddr, appConfig.kafkaTopic)
	if err != nil{
		logs.Error("init kafka failed, err:%v", err)
		return
	}


	// 初始化elasticsearch
	err = initES(appConfig.ESUrl)
	if err != nil{
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	// service start running
	err = run()
	if err != nil{
		logs.Error("transfer run failed, err:%v", err)
		return
	}
	
}