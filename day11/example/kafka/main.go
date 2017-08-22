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

	client, err := sarama.NewSyncProducer([]string{"192.168.14.7:9092"}, config)
	if err != nil {
		fmt.Println("kafka conn error:", err)
		return
	}
	defer client.Close()

	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("this is a good test, my message is good.")
		pid, offset, err := client.SendMessage(msg)

		if err != nil {
			fmt.Println("msg send err:", err)
			return
		}
		fmt.Printf("pid:%v, offset:%v\n", pid, offset)
		time.Sleep(time.Second)
	}

}
