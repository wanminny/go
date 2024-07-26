package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool) {

	fmt.Println("doing work!")
	time.Sleep(time.Second)
	ch <- true
}
func main() {

	done := make(chan bool, 1)
	go worker(done)
	<-done

	fmt.Println("finish job")
}
