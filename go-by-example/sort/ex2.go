package main

import (
	"log"
	"slices"
	"sort"
)

func main() {

	log.Println(sort.SearchInts([]int{
		1, 88, 99, 100,
	}, 99))

	a := []int{
		1, 3, 41, 41, 6,
	}
	log.Println(a)

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	log.Println(a)

	i, b := slices.BinarySearch(a, 6)
	log.Println(i, b)
}
