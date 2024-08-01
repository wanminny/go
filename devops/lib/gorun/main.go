package main

import (
	"context"
	"fmt"
	"github.com/oklog/run"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func test() {
	ctxAll, callcel := context.WithCancel(context.Background())
	log.Println(callcel)
	var g run.Group

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan struct{})
	{

		g.Add(func() error {
			log.Println("do work!")

			select {
			case sig := <-term:
				log.Println("rec signal", sig)
				callcel()
				//done <- struct{}{}
			case <-ctxAll.Done():
				log.Println("canceled ！")
			case <-done:
				log.Println("done!")
			}
			return nil
		}, func(err error) {
			log.Println("module 00  end")
			log.Println("打扫工作")
			close(done)
		})
	}

	{
		ticker := time.NewTicker(1 * time.Second)
		g.Add(func() error {
			for {
				select {
				case <-ctxAll.Done():
					log.Println("msg", "我是模块01退出了，接收到了cancelall")
					return nil
				case <-ticker.C:
					log.Println("msg", "我是模块01")
				}
			}
		}, func(err error) {
			if err != nil {
				log.Println("----", err.Error())
			}
			ticker.Stop()
			log.Println("module 01  end")
		})
	}
	{
		g.Add(func() error {
			for {
				select {
				case <-ctxAll.Done():
					log.Println("msg", "我是模块02退出了，接收到了cancelall")
					return nil
				}
			}
		}, func(err error) {
			if err != nil {
				log.Println("----", err.Error())
			}
			log.Println("module 02  end")
		})
	}
	err := g.Run()
	if err != nil {
		log.Println("err : ", err.Error())
	}
	log.Println("len(err):", err)
	//time.Sleep(time.Second)
	log.Println("end")

}

func test2() {
	var g run.Group

	// 第一个 goroutine: 模拟一个工作进程
	{
		g.Add(func() error {
			fmt.Println("工作进程开始运行")
			time.Sleep(3 * time.Second) // 模拟工作
			fmt.Println("工作进程运行结束")
			return nil
		}, func(error) {
			fmt.Println("m1 : 工作进程停止")
		})
	}

	// 第二个 goroutine: 监听信号以优雅地关闭
	{
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		g.Add(func() error {
			select {
			case s := <-sig:
				fmt.Printf("接收到信号: %v\n", s)
				return nil
			}
		}, func(error) {
			// 啥时候尽量的
			println("m2 interrupt")
			close(sig)
		})
	}

	// 运行所有 goroutine
	if err := g.Run(); err != nil {
		fmt.Printf("程序运行出错: %v\n", err)
	}

	fmt.Println("所有 goroutine 已停止")
}
func main() {
	//test()
	e1 := make(chan error, 2)
	log.Println(len(e1), e1)
	e1 <- nil // 空也是一个长度
	log.Println(len(e1), e1)

	test2()
}
