package main

import (
	"fmt"
	"time"
)

var notify = make(chan struct{})

func main() {

	fmt.Println("start")

	go func() {
		// 模拟做一下事情
		time.Sleep(time.Second)
		println("send notify ")
		notify <- struct{}{}
	}()

	<-notify

	fmt.Println("end")
}
