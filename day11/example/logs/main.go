package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})

	config["filename"] = "my.log"
	config["level"] = logs.LevelDebug
	fmt.Printf("%v\n", config)

	configStr, err := json.Marshal(config)

	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.Debug("this is debug, my name is %s", "stu01")
	logs.Trace("this is trace, my name is %s", "stu01")
	logs.Info("this is info, my name is %s", "stu01")
	logs.Warn("this is warning, my name is %s", "stu01")
	logs.Error("this is error, my name is %s", "stu01")
}
