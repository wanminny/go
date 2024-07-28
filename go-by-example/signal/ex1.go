package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("rev")
				time.Sleep(time.Second * 2)
				fmt.Println("clear up is over !")
				return
				//os.Exit(0)
			default:
				fmt.Println("server is running !")
				time.Sleep(time.Second)
			}
		}
	}()
	fmt.Println("wait....")
	//<-ctx.Done() // 这个不是必须的;
	fmt.Println("wait child to clear up!")
	wg.Wait()
	fmt.Println("end")
}
