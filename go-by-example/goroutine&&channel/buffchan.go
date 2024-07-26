package main

import (
	"fmt"
	"log"
	"sync"
)

var buffer = make(chan string, 2)

func bufferchan() {

	buffer <- "go"
	buffer <- "by example"
	close(buffer)
}
func main() {
	bufferchan()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range buffer {
			log.Println(v)
		}
	}()
	wg.Wait()
	fmt.Println("end")
}
