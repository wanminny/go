package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Event struct {
	cond *sync.Cond
	done bool
}

func NewEvent() *Event {
	return &Event{
		cond: sync.NewCond(&sync.Mutex{}),
		done: false,
	}
}

func (e *Event) Wait() {
	e.cond.L.Lock()
	defer e.cond.L.Unlock()

	/*

		for !e.done vs if !e.done :

		for !e.done：确保条件变量在被唤醒时再次检查条件。
		这种方式能够处理虚假唤醒（spurious wakeups）和避免丢失信号。

		if !e.done：仅在第一次检查条件。如果条件在第一次不满足，Wait 会阻塞当前 goroutine。永久阻塞了；
		因为错过了信号！！
		但是，如果条件变量被错误地唤醒或者信号在 Wait 被阻塞之前已经发出，if 语句将无法正确处理这些情况。

			为什么使用 for !e.done
			处理虚假唤醒（Spurious Wakeups）：
			虚假唤醒指的是等待中的 goroutine 在没有收到任何信号的情况下被唤醒。
			虽然在 Go 中这种情况不常见，但在一些系统中可能会发生。因此，使用 for 可以确保在虚假唤醒后继续等待。

			避免丢失信号：如果使用 if，而信号在 Wait 调用之前已经发出，等待的 goroutine 可能会永远阻塞，
			因为它错过了信号。使用 for 可以避免这种情况，因为条件会在每次唤醒后重新检查。
	*/
	for !e.done { // 用if 还是for 了？？
		log.Println("wait....")
		e.cond.Wait() // 这个不能一直阻塞吗？？
		log.Println("wait end")
	}
	fmt.Println("event occurred!")
}

func (e *Event) Notify() {

	e.cond.L.Lock()
	defer e.cond.L.Unlock()

	e.done = true
	//e.cond.Signal() // 是否可以用广播了？？
	e.cond.Broadcast()
}
func main() {

	//sync.WaitGroup{}.Wait()
	//sync.Cond{}
	ev := NewEvent()
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		ev.Wait()
	}()

	// 触发事件;
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 4)
		ev.Notify()
	}()
	wg.Wait()
}
