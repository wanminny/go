package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 0
	// 关闭与否是否有影响 ？？
	//close(ch)

	go func() {

		for {

			v, ok := <-ch
			if ok {
				log.Printf("is not closed,value is %d", v)
			} else {
				log.Printf("closed?  %t,%q", ok, v)
				time.Sleep(time.Second)
			}
			log.Println("one by one")
		}
	}()

	m := make(map[string]int)
	m["go"] = 1
	m["k"] = 0

	v, exists := m["k"]
	if exists {
		log.Println("exist.....", v)
	} else {
		log.Println("not exists... ", v)
	}

	//for {
	//	for k, v := range m {
	//		log.Println(k, v)
	//		time.Sleep(time.Second)
	//	}
	//}
	time.Sleep(time.Second * 3)
	log.Println("end.....")
}
