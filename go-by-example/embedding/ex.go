package main

import (
	"fmt"
	"log"
)

type base struct {
	//b() int
	salary float64
}

type container struct {
	base
	annotation string
}

func (b *base) Note() {
	fmt.Println("base ", b.salary)
}

func main() {
	c := &container{
		base{
			30000,
		},
		"aaaaaaaa",
	}
	log.Println(c.salary)
	log.Println(c.base.salary, c.annotation)
	c.Note()

	type career interface {
		Note()
	}
	var ee career = c
	ee.Note()
}
