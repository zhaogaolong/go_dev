package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func process() {

	ctx, cancle := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancle()

	tr := &http.Transport{}

	client := &http.Client{Transport: tr}

	c := make(chan Result, 1)

	res, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}

	go func() {

		resp, err := client.Do(res)
		pack := Result{r: resp, err: err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(res)
		res := <-c
		fmt.Println("time out, err:", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Println("server Respone:", string(out))
	}
	return

}

func main() {
	process()
}
