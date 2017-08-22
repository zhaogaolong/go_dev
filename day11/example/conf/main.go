package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "my.conf")
	if err != nil {
		fmt.Println("open config failed, err:", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("get port failed,err:", err)
		fmt.Println("must set default config port")
		port = 8080
	}
	fmt.Println("port:", port)

	log_level := conf.String("log::level")
	fmt.Println("log level:", log_level)
}
