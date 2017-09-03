package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

func initLogger(logLevel, logPath string)(err error){

	logConf := make(map[string]interface{})
	logConf["filename"] = logPath
	switch logLevel {
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
	logs.Debug("logger init success.")
	return
}