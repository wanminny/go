package main

import (
	"fmt"
	"log"
	"sync"
)

func even(i int) bool {
	return i%2 != 0
}

func odd(in int) bool {
	return in%2 == 0
}

type fn func(int) bool

func filter(c []int, f fn) []int {
	var rlt []int
	for _, v := range c {
		if f(v) {
			rlt = append(rlt, v)
		}
	}
	return rlt
}

func sum(a int, plus ...int) int {
	var sum = 0
	for _, v := range plus {
		sum += v
	}
	return sum + a
}

func clourse() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		tmp := i
		go func() {
			defer wg.Done()
			fmt.Println(tmp)
		}()
	}
}

func req() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

var wg sync.WaitGroup

func main() {
	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	log.Println(filter(input, even))
	log.Println(filter(input, odd))

	log.Println(sum(1, 2, 3, 4))
	var input1 = []int{10, 40}
	log.Println(sum(1, input1...))

	clourse()
	wg.Wait()

	f1 := req()
	fmt.Println(f1())
	fmt.Println(f1())
	
}
