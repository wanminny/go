package main

import "log"

type Element[T any] struct {
	v    T
	next *Element[T]
}

type List[T any] struct {
	head, tail *Element[T]
}

func (l *List[T]) Push(v T) {
	tail := l.tail
	if tail == nil {
		e := &Element[T]{v: v}
		l.tail = e
		l.head = l.tail
	} else {
		e := &Element[T]{v: v}
		l.tail.next = e
		l.tail = l.tail.next
	}
}

func (l *List[T]) GetAll() []T {
	var rlt []T
	head := l.head
	for head != nil {
		v := head.v
		rlt = append(rlt, v)
		head = head.next
	}
	return rlt
}

func KeyOfMaps[K comparable, V any](m map[K]V) []K {
	rlt := make([]K, 0, len(m))
	for k, _ := range m {
		rlt = append(rlt, k)
	}
	return rlt
}
func main() {
	list1 := List[int]{}
	list1.Push(1)
	list1.Push(2)
	list1.Push(99)
	log.Println(list1)
	log.Println(list1.GetAll())

	log.Println("---------")
	var m2 = make(map[string]string)
	m2["go"] = "100"
	m2["by"] = "200"
	m2["example"] = "300"
	m1 := KeyOfMaps[string, string](m2)
	m3 := KeyOfMaps(m2)

	log.Println(m1, m3)
}
