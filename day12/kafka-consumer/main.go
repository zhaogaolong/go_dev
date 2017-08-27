package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func main() {

	consumer, err := sarama.NewConsumer(strings.Split("192.168.14.4:9092", ","), nil)

	if err != nil {
		fmt.Println("failed, err:", err)
		return
	}

	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Println("failed get partitions for nginx_log topic.")
		return
	}

	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("failed to start consumer")
			return
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {

			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
				fmt.Println()
			}
		}(pc)
	}
	time.Sleep(time.Hour)
	consumer.Close()
}
