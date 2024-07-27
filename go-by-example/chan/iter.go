package main

import "log"

func main() {
	queue := make(chan string, 2)
	queue <- "go"
	queue <- "by"

	close(queue)

	//方法一
	//for v := range queue {
	//	log.Println(v)
	//}
	// 方法二
	for {
		msg, ok := <-queue
		if ok {
			log.Println(msg)
		} else {
			log.Println("closed")
			return
		}
	}
}
