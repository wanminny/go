package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal)
	//sigs := make(chan os.Signal,1) //是否有缓存均可以！
	done := make(chan struct{})

	// os 或者 syscall包中的信号
	//signal.NotifyContext()
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, os.Kill)

	go func() {
		sig := <-sigs
		log.Println("rec: ", sig)
		done <- struct{}{}
	}()
	fmt.Println("wait signal of exit!")
	<-done
	fmt.Println("end")

}
