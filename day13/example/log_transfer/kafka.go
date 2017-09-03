package main

import (
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

type KafkaClient struct{
	client sarama.Consumer
	addr string
	topic string
	wg sync.WaitGroup
}


var (
	kafkaClient *KafkaClient
)



func initKafka(addr , topic string) (err error){

	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)

	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	kafkaClient = &KafkaClient{}
	kafkaClient.client = consumer
	kafkaClient.addr = addr
	kafkaClient.topic = topic
	logs.Info("kafka init success")
	return

	/*
	//consumer.Close()
	return
	*/
}
