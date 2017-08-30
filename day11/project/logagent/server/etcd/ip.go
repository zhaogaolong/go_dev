package etcd

import (
	"net"

	"github.com/astaxie/beego/logs"
)

var localIPArry []string

func init() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		logs.Error("Can't get host ipAddress")
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPArry = append(localIPArry, ipnet.IP.String())
			}
		}
	}
	// fmt.Println(localIPArry)
}
