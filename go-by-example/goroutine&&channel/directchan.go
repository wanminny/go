package main

import (
	"fmt"
	"sync"
)

func ping(in chan<- string, msg string) {
	//defer wg.Done()
	//fmt.Println("ping .....")
	in <- msg
}

func pong(in <-chan string, msg chan<- string) {
	//defer wg.Done()

	//msg <-  <- in
	rec := <-in
	fmt.Println(rec)
	msg <- "pong"
}

var wg = sync.WaitGroup{}

func main() {

	pingMsg := make(chan string)
	pongMsg := make(chan string)
	//wg.Add(2)
	go ping(pingMsg, "ping")
	go pong(pingMsg, pongMsg)

	fmt.Println(<-pongMsg)
	wg.Wait()
}
