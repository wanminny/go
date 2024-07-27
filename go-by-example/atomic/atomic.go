package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var count uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				atomic.AddUint64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println(atomic.LoadUint64(&count))
}
