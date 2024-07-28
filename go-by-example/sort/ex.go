package main

import (
	"fmt"
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

	example := []int{1, 25, 3, 5, 4}
	// IntSlice是类型；这个是转换不是函数！！！
	sort.Sort(sort.Reverse(sort.IntSlice(example)))
	fmt.Println(example)

	//sort.IntSlice(intS)

}
