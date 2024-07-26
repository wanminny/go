package main

import (
	"fmt"
	"time"
)

type compare func(a, b int) bool

func cmp1(a, b int) bool {
	return a > b
}

func sort(input []int, c compare) {
	fmt.Println("func sort.....")
	c(input[0], input[1])

}

func sort1(input []int, r Revert) {
	r.cmp(input[0], input[1])
}

type Revert interface {
	cmp(a, b int) bool
}

type ACmp struct {
}

func (ins *ACmp) cmp(a, b int) bool {
	fmt.Println("interface sort.....")
	return a > b
}

var setup func()
var tearDown func()

func init() {
	tearDown = func() {
		fmt.Println("收尾工作")
	}
}
func main() {
	setup = func() {
		fmt.Println("初始化的工作")
	}
	setup()
	time.Sleep(time.Second)
	tearDown()

	s := []int{1, 2}
	sort(s, cmp1)

	acmp := &ACmp{}
	sort1(s, acmp)
}
