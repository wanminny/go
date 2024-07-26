package main

import "log"

func swap(a, b int) {
	a, b = b, a
}

func swap1(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	a := 1
	b := 2
	log.Println(a, b)
	swap(a, b)
	log.Println(a, b)

	swap1(&a, &b)
	log.Println(a, b)

}
