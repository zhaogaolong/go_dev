package main

import (
	"fmt"

	elastic "gopkg.in/olivere/elastic.v2"
)

type Tweet struct {
	User    string
	Massage string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.1.1:9200"))
	if err != nil {
		fmt.Println("conn failed.")
		return
	}

	fmt.Println("conn succ")
	tweet := Tweet{User: "oliver", Massage: "Task live"}
	_, err = client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do()

	if err != nil {
		panic(err)
		return
	}

}
