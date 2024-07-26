package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	mySignal := make(chan os.Signal, 1)
	signal.Notify(mySignal, os.Interrupt, os.Kill)
	log.Println("start .....")
	time.Sleep(time.Second)
	select {
	case <-ch1:
		fmt.Println("取到了ch1的消息")
	default:
		fmt.Println("没有收到消息吖。不能阻塞")
	}

	msg := "hi"
	select {
	case ch1 <- msg:
		fmt.Println("ch1 接受了消息")
	default:
		fmt.Println("ch1 还是阻塞了，所以来到这里")
	}
	log.Println("recv signal  .....")
	time.Sleep(time.Second)
	select {
	case s := <-mySignal:
		fmt.Println("收到了 信号", s)
	case rev := <-ch1:
		fmt.Println(rev)
	default:
		fmt.Println("默认分支。。。")
	}

}
