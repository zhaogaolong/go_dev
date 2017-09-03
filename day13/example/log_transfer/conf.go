package main


import (
	"fmt"
	"github.com/astaxie/beego/config"
)
type LogConfig struct{
	kafkaAddr string
	kafkaTopic string
	ESUrl string
	LogPath string
	LogLevel string
}

var (
	appConfig *LogConfig
)

func initConfig(configType, fileName string)(err error){
	conf, err := config.NewConfig(configType, fileName)
	if err != nil {
		err = fmt.Errorf("open config failed, err:%v", err)
		return
	}
	appConfig = &LogConfig{}

	// init logger
	appConfig.LogLevel = conf.String("log::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "info"
	}
	appConfig.LogPath = conf.String("log::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogLevel = "./logagent.log"
	}

	// init config for kafka
	appConfig.kafkaAddr = conf.String("kafka::addr")
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}
    appConfig.kafkaTopic = conf.String("kafka::topic")
    if len(appConfig.kafkaTopic) == 0 {
        err = fmt.Errorf("invalid kafka topic")
        return
	}
	
	// init config for es
	appConfig.ESUrl = conf.String("es::url")
	if len(appConfig.ESUrl) == 0 {
		err = fmt.Errorf("invalid es url")
		return
	}
	return
}