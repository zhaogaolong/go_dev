package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"192.168.1.201:9092"}, config)
	if err != nil {
		fmt.Println("kafka conn error:", err)
		return
	}
	defer client.Close()
	var count int
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("my name is Tony, 一起打王者荣耀." + fmt.Sprintf("%d", count))
		pid, offset, err := client.SendMessage(msg)

		if err != nil {
			fmt.Println("msg send err:", err)
			return
		}
		fmt.Printf("pid:%v, offset:%v\n", pid, offset)
		time.Sleep(time.Second)
		count++
	}

}
