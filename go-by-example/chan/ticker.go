package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func ticker() {
	wg := sync.WaitGroup{}
	tk1 := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
	LABEL:
		for {
			select {
			case <-done:
				fmt.Println("done!")
				break LABEL // 这样可以效果等于 return
				//return
				//break  // break 只是针对select所以要用return

			case <-tk1.C:
				log.Println("dida dida ....")
			}
		}
	}()
	time.Sleep(time.Millisecond * 1600)
	tk1.Stop()
	done <- true
	wg.Wait()
	//time.Sleep(time.Millisecond)
	fmt.Println("end")
}
func main() {
	ticker()
}
