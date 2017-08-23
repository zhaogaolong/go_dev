package kafka

import (
	"fmt"
	"go_dev/day11/project/logagent/server"
	"go_dev/day11/project/logagent/server/tailf"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	kafkaClient sarama.SyncProducer
)

func ServerRun(addr string) {
	// 连续检查初始化kafka连接
	for {
		err := initKafka(addr)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		defer kafkaClient.Close()
		logs.Debug("kafka client init success")
		break
	}

	for {
		msg := <-tailf.TailfObjsMgr.MsgChan

		// fmt.Println("get from chan msg:", msg)
		sendToKafka(msg)
	}

}

// 初始化kafka连接
func initKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	kafkaClient, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		err = fmt.Errorf("kafka can't init, err:%v", err)
	}
	return
}

// 发送数据到kafka
func sendToKafka(message *server.TextMsg) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = message.Topic
	msg.Value = sarama.StringEncoder(message.Msg)

	// pid, offset, err := kafkaClient.SendMessage(msg)
	_, _, err := kafkaClient.SendMessage(msg)
	if err != nil {
		logs.Error("send massage failed err:%v", err)
		return
	}
}
