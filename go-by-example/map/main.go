package main

import "log"

func main() {

	var m1 map[string]int
	log.Println(m1, len(m1))

	m2 := make(map[string]map[string]int, 2)
	m2["go"] = make(map[string]int, 2)
	m2["by"] = make(map[string]int, 1)

	log.Println(m2)
	m2["go"]["wm"] = 1
	m2["go"]["njupt"] = 2
	m2["go"]["cs"] = 3

	log.Println(m2, len(m2), len(m2["go"]))

	m3 := make(map[int]int)

	for i := 0; i < 3; i++ {
		for j := 1; j < 2; j++ {
			m3[i] = j
		}
	}
	log.Println(m3)

	for k, v := range m3 {
		log.Println(k, v)
	}

	for k, v := range "go" {
		log.Println(k, v)
	}

}
