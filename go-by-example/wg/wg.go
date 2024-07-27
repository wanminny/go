package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("start worker %d\n", id)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("end worker %d\n", id)

}
func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go worker(i, &wg)

		// 下面也可以
		//go func() {
		//	defer wg.Done()
		//	worker(i)
		//}()
	}
	wg.Wait()

}
