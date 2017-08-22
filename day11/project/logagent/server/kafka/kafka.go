package kafka

import (
	"go_dev/day11/project/logagent/server"
	"go_dev/day11/project/logagent/server/tailf"

	"github.com/astaxie/beego/logs"
)

func SendToKafka(msg *server.TextMsg) {

	msg = <-tailf.TailfObjsMgr.MsgChan
	logs.Debug("read msg:%s, topic:%s,", msg.Msg, msg.Topic)

}
