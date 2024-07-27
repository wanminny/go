package main

import (
	"fmt"
	"log"
	"time"
)

const LAYOUT = "2006-01-02 15:04:05"

func test1() {
	log.Println(time.Now().Format(LAYOUT))
	timer1 := time.NewTimer(time.Second * 2)
	defer timer1.Stop()
	// 一次性的会阻塞死锁
	for {
		select {
		case msg := <-timer1.C:
			log.Println(msg.Format(LAYOUT))
			timer1.Stop()
		}
	}
}

func test2() {
	log.Println(time.Now().Format(LAYOUT))
	timer1 := time.NewTicker(time.Second * 2)

	for {
		// 周期性的不会死锁
		select {
		case msg := <-timer1.C:
			log.Println(msg.Format(LAYOUT))
		}
	}
}
func test3() {
	timer3 := time.After(time.Second * 2)
	//defer
	//for {
	// 也是一次性的
	select {
	case msg := <-timer3:
		log.Println(msg.Format(LAYOUT))
	}
	//}
}

func test4() {
	timer4 := time.NewTimer(time.Second)
	<-timer4.C
	fmt.Println("timer 4 fired")
	defer timer4.Stop()

	//timer5 := time.NewTimer(time.Second * 2)

	var timer5 *time.Timer
	go func() {
		timer5 = time.NewTimer(time.Second * 2)
		<-timer5.C
		fmt.Println("timer 5 fired")
	}()
	time.Sleep(time.Millisecond * 10)
	if timer5.Stop() {
		log.Println("time stopped")
	}
	fmt.Println("end")
	time.Sleep(time.Second * 3)
}
func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
