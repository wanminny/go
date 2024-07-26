package main

import "log"

type rectangle struct {
	Len, Width int
}

func (r rectangle) area() int {
	return r.Len * r.Width
}

func (r rectangle) circle() int {
	return (r.Len + r.Width) << 1
}
func main() {
	r := rectangle{}
	r.Len = 5
	r.Width = 4

	log.Println(r.area(), r.circle())

	r1 := &rectangle{
		5,
		4,
	}
	log.Println(r1.area(), r1.circle())
}
