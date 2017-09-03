package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

func run()(err error){
	partitionList, err := kafkaClient.client.Partitions(kafkaClient.topic)
	if err != nil {
		logs.Error("failed get partitions for nginx_log topic.")
		return
	}

	logs.Debug("start transfer running...")

	for partition := range partitionList {
		pc, err := kafkaClient.client.ConsumePartition(
			kafkaClient.topic, 
			int32(partition), 
			sarama.OffsetNewest,
		)
		if err != nil {
			continue
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			kafkaClient.wg.Add(1)
			for msg := range pc.Messages() {
				logs.Info(
					"Partition:%d, Offset:%d, Key:%s, Value:%s\n", 
					msg.Partition, msg.Offset, msg.Key, msg.Value,
				)
				err = sendToES(kafkaClient.topic,msg.Value)
				if err != nil{
					logs.Error("send message to es failed, err:%v", err)
				}
			}
			kafkaClient.wg.Done()
		}(pc)
	}
	kafkaClient.wg.Wait()

	return
}

