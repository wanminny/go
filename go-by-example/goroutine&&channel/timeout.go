package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "c1.."
	}()
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	select {

	case <-c1:
		fmt.Println("c1")
	case <-timer.C:
		fmt.Println("c1超时了。。")
	}
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 5)
		c2 <- "c2.."
	}()
	timer1 := time.NewTimer(time.Second * 6)
	defer timer1.Stop()
	select {
	case <-c2:
		fmt.Println("c2")
	case <-timer1.C:
		fmt.Println("c1超时了。。")
	}
}
