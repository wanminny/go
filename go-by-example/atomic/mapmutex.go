package main

import (
	"log"
	"sync"
)

type Container struct {
	sync.Mutex
	counters map[string]int
}

func (c *Container) Inc(name string) {
	c.Lock()
	defer c.Unlock()
	c.counters[name]++
}

func main() {
	wg := sync.WaitGroup{}
	c1 := Container{
		counters: map[string]int{
			"a": 0,
			"b": 1,
		},
	}
	doSth := func(name string, c int) {
		defer wg.Done()
		for i := 0; i < c; i++ {
			c1.Inc(name)
		}
	}
	wg.Add(3)
	go doSth("a", 10000)
	go doSth("b", 10000)
	go doSth("a", 10000)
	wg.Wait()
	log.Println(c1.counters)

}
