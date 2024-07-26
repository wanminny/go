package main

import (
	"fmt"
	"log"
	"sync"
)

func doSth(from string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("%s %d\n ", from, i)
	}
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	doSth("no go routine", wg)

	go doSth("goroutine", wg)

	go func() {
		defer wg.Done()
		fmt.Println("匿名函数")
	}()
	wg.Wait()
	log.Println("end")

}
