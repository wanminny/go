package main

import (
	"fmt"
	"log"
	"time"
)

func tick() {
	t := time.Tick(time.Millisecond * 500)

	for {
		select {
		case <-t:
			fmt.Println("500 ms tick")
		}
	}
}
func main() {
	req := make(chan int, 5)
	for i := 0; i < 5; i++ {
		req <- i
	}
	close(req)
	limiter := time.Tick(time.Millisecond * 200)
	// req 200ms一个 限速
	for v := range req {
		<-limiter
		log.Printf("req %d\n", v)
	}

	// bursty 突发流量
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 800) {
			//log.Println("---------")
			burstyLimiter <- t
		}
	}()

	burstyReq := make(chan int, 10)
	for i := 0; i < 10; i++ {
		burstyReq <- i
	}
	close(burstyReq)

	//模拟 10个请求，前三个非常快;后面7个要ticker 限速！！！
	for req := range burstyReq {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
	//time.Sleep(time.Second * 3)
}
