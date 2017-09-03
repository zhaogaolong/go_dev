package main

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v2"
)


type Tweet struct{
	User string
	Message string
}

func main(){
	client, err := elastic.NewClient(
		elastic.SetSniff(false), 
		elastic.SetURL("http://192.168.1.201:9200/"),
	)
	if err != nil{
		fmt.Println("elastic connect failed, err:", err)
		return
	}
	fmt.Println("elastic connect success")
	for i := 0; i < 100; i++{
		tweet := Tweet{User:"olivere", Message:"Take Five"}
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id(fmt.Sprintf("%d", i)).
			BodyJson(tweet).
			Do()

		if err != nil{
			panic(err)
			return
		}
		
	}
	fmt.Println("inset success")
}
