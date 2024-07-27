package main

import (
	"fmt"
	"log"
	"time"
)

func closeChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			msg, ok := <-jobs
			if ok {
				fmt.Println(msg)
				time.Sleep(time.Millisecond * 500)

			} else {
				fmt.Println("no msg, closed")
				done <- true
				break
			}
		}
	}()

	for i := 0; i < 4; i++ {
		jobs <- i
	}
	//生产方关闭
	close(jobs)

	<-done
	fmt.Println("end")
}

func test() {
	var c = make(chan string, 3)
	c <- "go"
	c <- "by"
	//c <- "example"
	//go func() {
	// 一直遍历会死锁阻塞
	for v := range c {
		log.Println(v)
	}
	//}()
}
func main() {

	test()
	closeChannel()
	time.Sleep(time.Second)
}
