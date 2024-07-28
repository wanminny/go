package main

import (
	"log"
	"sort"
)

// 通过长度来排序
type SizeLen struct {
	content []string
}

func (s *SizeLen) Less(i, j int) bool {
	// 从小到大排序
	return len(s.content[i]) < len(s.content[j])
}

func (s *SizeLen) Swap(i, j int) {
	s.content[i], s.content[j] = s.content[j], s.content[i]
}

func (s *SizeLen) Len() int {
	return len(s.content)
}

func main() {

	size := &SizeLen{
		content: []string{"primary-is-biggest", "go-by-example", "data-structure-golang"},
	}
	// 排序前
	log.Println(size)
	log.Println(sort.IsSorted(size))

	sort.Sort(size)
	// 排序后
	log.Println(size)
	log.Println(sort.IsSorted(size))

	//逆向反转
	sort.Sort(sort.Reverse(size))
	log.Println(size)

}
