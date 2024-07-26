package main

import (
	"log"
	"slices"
	"sort"
)

func slice() {
	s1 := []int{}
	//s1[0] = 1
	s1 = append(s1, 2)
	s1 = append(s1, 66)
	log.Println(s1, len(s1), cap(s1))
	s1 = append(s1, 68)
	//s1 = append(s1, 66)
	log.Println(s1, len(s1), cap(s1))

	var s2 [][]int
	log.Println(s2)

	s3 := make([][]int, 2)
	log.Println(s3)
	s3[0] = make([]int, 1)
	log.Println(s3)
	s3[0][0] = 1
	//s3[0][1] = 2
	//s3[0][2] = 3
	s3[0] = append(s3[0], 2)
	log.Println(s3)

	s4 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		tmp := i + 1
		s4[i] = make([]int, tmp)
		for j := 0; j < tmp; j++ {
			s4[i][j] = j
		}
	}

	log.Println(s4)

	var s5 []int
	log.Println(len(s5), s5, cap(s5))

	s5 = append(s5, 100)
	log.Println(len(s5), s5, cap(s5))

	s6 := make([]int, len(s1))
	copy(s6, s1)
	log.Println(s6)
	//s6 = append(s6[0:1],
	log.Println(s6[:], s6[0:], s6[0:1], s6[len(s6)-1:])
	s6 = append(s6[0:1], s6[2:]...)
	log.Println(s6)
	//copy(s6[0:1], s6[1:])
	log.Println(s6)

}

func map1() {
	var m1 map[string]int
	//m1["a"] = 1
	log.Println(m1)
	m2 := make(map[string]int)
	m2["go"] = 1
	m2["by"] = 2
	m2["example"] = 3
	log.Println(m2)
	for k, v := range m2 {
		log.Println(k, v)
	}
	var m3key []string
	for k, _ := range m2 {
		m3key = append(m3key, k)
	}
	log.Println(m3key)
	sort.Strings(m3key)
	log.Println(m3key)
	for _, v := range m3key {
		log.Println(m2[v])
	}
	s1 := []int{4, 3}
	s2 := []int{8, 7}
	s3 := slices.Concat(s1, s2)
	log.Println(s3)

	delete(m2, "go")
	log.Println(len(m2))
	log.Println(m2)
}

func main() {

	//slice()
	map1()
}
