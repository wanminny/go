package main

import (
	"fmt"
	"math"
)

type geo interface {
	area() float64
	permi() float64
}

type rect struct {
	l, w float64
}
type circle struct {
	r float64
}

// 面向接口编程
func measure(g geo) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.permi())
}

func (r *rect) area() float64 {
	return r.w * r.l
}
func (r *rect) permi() float64 {
	return (r.l + r.w) * 2
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *circle) permi() float64 {
	return 2 * math.Pi * c.r
}
func main() {
	c := circle{}
	c.r = 100
	r := rect{
		w: 10,
		l: 12,
	}
	measure(&c)
	measure(&r)
}
