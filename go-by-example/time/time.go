package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	now.Unix()
	//log.Println(now.Unix())
	//log.Println(now.Nanosecond())
	//log.Println(now.UnixNano())
	sec := now.Unix()
	nanos := now.UnixNano()
	log.Println(now)
	log.Println(sec)
	log.Println(nanos)
	mills := nanos / 1000000
	log.Println(mills)

	fmt.Println(time.Unix(sec, 0))
	fmt.Println(time.Unix(0, nanos))
	log.Println("=================")
	fmt.Println(time.Unix(sec, 0))
	fmt.Println(time.Unix(0, nanos))
}
