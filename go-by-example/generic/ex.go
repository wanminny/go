package main

type Element[T any] struct {
	v    T
	next *Element[T]
}

type List[T any] struct {
	head, tail *Element[T]
}

func main() {

}
