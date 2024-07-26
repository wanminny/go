package main

import (
	"fmt"
	"log"
)

type Career struct {
	Name   string
	Salary float64
}

func (c *Career) Owns() {
	fmt.Println("name", c.Name, " salary", c.Salary)
}

type Student struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Addr string `json:"addr"`
	Career
	//c *Career
}

func NewStudent() *Student {
	// 要全部写完整否则爆红
	s := &Student{
		"ethan",
		10,
		"nj",
		Career{},
	}
	s.Age = 11
	//s->Age
	(*s).Age = 12
	//*s.Addr = "sh"
	return s
}

//func (s *Student) Owns() {
//	fmt.Println("student .....")
//}

func main() {
	//s := Student{
	//	Name: "tony",
	//	Age:  33,
	//}
	s := NewStudent()
	s.Addr = "nj-sh"
	log.Printf("%#v", s)
	//log.Println(s.Owns())
	//s.Owns()

	log.Println(NewStudent())
}
