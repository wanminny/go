package main

import (
	"log"
	"slices"
	"sort"
)

func main() {
	s := []string{"b", "a", "g", "c"}
	log.Println(s)
	//slices.Sort(s)
	//log.Println(s)

	sort.Strings(s)
	log.Println(s)

	intS := []int{5, 7, 8, 9, 0, 11, 3}
	log.Println(intS)
	// 也可以使用slices包排序
	//slices.Sort(intS)
	//log.Println(intS)

	//可以使用sort包排序
	sort.Ints(intS)
	log.Println(intS)

	//同理判断是否有序；也有两种方法 分别是使用slices和sort
	log.Println(slices.IsSorted(s), slices.IsSorted(intS))

	log.Println(sort.IntsAreSorted(intS), sort.StringsAreSorted(s))

}
