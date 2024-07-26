package main

import (
	"fmt"
	"time"
)

func fn1(in chan<- string) {
	fmt.Println("fn1.函数...")
	time.Sleep(time.Second)
	in <- "fn1"
	//close(in)
}

func fn2(in chan<- string) {
	fmt.Println("fn2.. 函数...")
	time.Sleep(time.Second)
	in <- "fn2"
	//close(in)
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	go fn1(ch1)
	go fn2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
			time.Sleep(time.Second)

		case msg2 := <-ch2:
			fmt.Println(msg2)
			time.Sleep(time.Second)
		}
	}
	//for {
	//
	//	select {
	//	case msg1 := <-ch1:
	//		fmt.Println(msg1)
	//		time.Sleep(time.Second)
	//
	//	case msg2 := <-ch2:
	//		fmt.Println(msg2)
	//		time.Sleep(time.Second)
	//
	//		//default:
	//		//	time.Sleep(time.Second)
	//		//	fmt.Println("default...")
	//	}
	//}
}
