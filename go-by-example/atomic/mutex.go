package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"sync"
)

// 因为要有++运算符所以要自定义泛型
type Number interface {
	constraints.Integer | constraints.Float
}

// ConNum 并发安全的数字 添加泛型
type ConNum[T Number] struct {
	num T
	sync.Mutex
}

func (c *ConNum[T]) Inc() {
	c.Lock()
	defer c.Unlock()
	c.num++
}
func (c *ConNum[T]) GetNum() T {
	c.Lock()
	c.Unlock()
	return c.num
}

func main() {
	//cN := new(ConNum[uint64])
	cN := ConNum[float64]{
		num: 1.0,
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				cN.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("%.2f\n", cN.GetNum())
	fmt.Println(cN.num)
	fmt.Printf("%.2f", cN.num)

}
