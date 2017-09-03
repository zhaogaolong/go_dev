package main

import (
	"github.com/astaxie/beego/logs"
	elastic "gopkg.in/olivere/elastic.v2"
)


type LogMessage struct{
	App string
	Topic string
	Message string
}


var (
	esClient *elastic.Client
)

func initES(url string)(err error){
	client, err := elastic.NewClient(
		elastic.SetSniff(false), 
		elastic.SetURL(url),
	)
	if err != nil{
		logs.Error("elastic connect failed, err:", err)
		return
	}
	esClient = client
	logs.Info("elasticsearch init success")
	return
}

func sendToES(topic  string,message []byte)(err error){
	msg := &LogMessage{}
	msg.App = topic
	msg.Topic = topic
	msg.Message = string(message)


	_, err = esClient.Index().
		Index(topic).
		Type(topic).
		BodyJson(msg).
		Do()

	if err != nil{
		panic(err)
		return
	}

	return
}