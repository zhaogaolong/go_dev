package main

import (
	"encoding/json"
	"fmt"

	"go_dev/day11/project/logagent/server"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

// config logagent config
// collectConf save some collect log config
type Config struct {
	logLevel    string
	logPath     string
	collectConf []server.CollectConf
	chanSize    int
	kafka_addr  string
}

var (
	appConf *Config
)

func loadConf(configType, fileName string) (err error) {
	conf, err := config.NewConfig(configType, fileName)
	if err != nil {
		fmt.Println("open config failed, err:", err)
		return
	}
	appConf = &Config{}
	loadLogsConfig(conf)
	loadKafkaConfig(conf)
	err = loadColletConf(conf)
	if err != nil {
		logs.Warn("load colect conf failed, err:%v", err)
		return
	}
	return
}

// init logger config
func loadLogsConfig(conf config.Configer) {
	appConf.logLevel = conf.String("log::log_level")
	if len(appConf.logLevel) == 0 {
		appConf.logLevel = "info"
	}

	appConf.logPath = conf.String("log::log_path")
	if len(appConf.logLevel) == 0 {
		appConf.logLevel = "./logagent.log"
	}

	logConf := make(map[string]interface{})
	logConf["filename"] = appConf.logPath
	switch appConf.logLevel {
	case "info":
		logConf["level"] = logs.LevelInfo
	case "debug":
		logConf["level"] = logs.LevelDebug
	default:
		logConf["level"] = logs.LevelInfo
	}

	logConfStr, err := json.Marshal(logConf)
	if err != nil {
		logs.Error("Can't marshal log map")
		return
	}

	logs.SetLogger(logs.AdapterFile, string(logConfStr))
	logs.Debug("set logs config success.")

}

func loadKafkaConfig(conf config.Configer) {
	appConf.kafka_addr = conf.String("kafka::server_addr")
	if len(appConf.kafka_addr) == 0 {
		appConf.logLevel = "localhost:9092"
		logs.Debug("set default config kafka_addr=%s", appConf.kafka_addr)
	}

}
func loadColletConf(conf config.Configer) (err error) {

	appConf.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		size := 100
		appConf.chanSize = size
		logs.Warn("collect chan_size can't get, must config defule size:%v", size)
		return
	}

	var cc server.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	cc.Topic = conf.String("collect::topic")

	if len(cc.LogPath) == 0 {
		err = fmt.Errorf("invalid [conllect]::log_path")
		return
	}

	if len(cc.Topic) == 0 {
		err = fmt.Errorf("invalid [conllect]::log_topic")
		return
	}

	appConf.collectConf = append(appConf.collectConf, cc)
	return
}
