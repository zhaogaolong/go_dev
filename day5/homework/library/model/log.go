package model

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

var (
	log    = logging.MustGetLogger("example")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc}  %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	LogFile, openErr = os.OpenFile("library.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	LogFileOpened    bool
)

func init() {

	if openErr != nil {
		fmt.Println("Can't open library.log file")
	}

	// 重置log文件打开状态，方便程序退出后关闭该文件句柄
	LogFileOpened = true

	backend2 := logging.NewLogBackend(LogFile, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend2Leveled := logging.AddModuleLevel(backend2)
	backend2Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend2Leveled, backend2Formatter)

	// log.Debugf("debug %s", "*****")
	// log.Info("info")
	// log.Notice("notice")
	// log.Warning("warning")
	// log.Error("err")
	// log.Critical("crit")

}
