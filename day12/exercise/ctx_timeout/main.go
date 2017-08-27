package main

import (
	"io/ioutil"
	"time"
	"net/http"
	"context"
	"fmt"
)

type Result struct{
	r *http.Response
	err error
}

func process(){
	// 生成一个context超时多对象
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// 最后关闭这个对象
	defer cancel()

	//生成http的相关对象，和返回结果的管道
	tr := &http.Transport{}
	client :=  &http.Client{Transport: tr}
	resultChan := make(chan Result, 1)

	// 生成一个http请求的对象
	req, err := http.NewRequest("GET", "https://www.google.com.hk", nil)
	if err != nil{
		fmt.Println("http request failed, err:", err)
		return
	}

	// 使用goroutine 执行开始请求并把结果返回到resultChan通道中
	go func(){
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		resultChan <- pack
	}()

	// 检查两个通道判断请求的结果 
	// ctx.Done()是一个超时的通道，如果之行后就会在指定的超时时间点放入一个消息到这个通道中
	// resultChan 是请求结果的通道，如果这个通道有数据说明请求结果完成。
	select {
	case <- ctx.Done():
		// 是否能在超时时间通道获取数据，如果获取到就直接把该请求取消。
		tr.CancelRequest(req)
		res := <- resultChan
		fmt.Println("Time out, err:", res.err)
	case res := <- resultChan :
		// 如果有数据打印请求的数据
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("server Respones :%s", out)
	}


}

func main(){
	process()
}