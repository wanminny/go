package main

import (
	"fmt"
	"log"
	"time"
)

// jobs 里面只能取
// result是存放结果的只能存入
// workId 做为worker的一个区分；看有多少个worker在处理的明细
func worker(jobs <-chan int, result chan<- int, workId int) {
	for v := range jobs {
		// 取出来jobs 来计算后将结果存入result
		// 模拟取出Job后的计算处理动作
		fmt.Printf("start worker id : %d on job %d\n", workId, v)
		time.Sleep(time.Millisecond * 500)
		//fmt.Printf("end worker id : %d on job %d\n", workId, v)

		//result <- 2 * v
		result <- v
	}

	//// jobs处理完成后要关闭 扇出的通道 不能这里关闭会 有可能其他worker还在往里面发送数据
	//close(result)
}

// 这个其实是一个经典的扇入扇出程序； jobs 扇入 三个 工作池协程 ；
// 然后处理后扇出到一个result通道。在主协程中处理
func main() {

	/// 多少个jobs 要处理; 所有的job是通过一个 有缓存的通道放入进去的；
	// 如果个 worker共享这个job 通道的；即所有工作 goroutine 都是从里面取job来计算的；
	// 计算处理后的结果放入到一个 通道中去;
	// 最后主协程 只消费这个结果通道;

	// 需要处理的jobs个数; 用通道存进去
	const numJobs = 10
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	// 需要启动多少个工作池 协程
	// 更具实际情况 开启格式
	const workerNums = 3
	for i := 0; i < workerNums; i++ {
		go worker(jobs, result, i+1)
	}

	// 放入真实要处理的jobs；jobs生产者
	for i := 0; i < numJobs; i++ {
		jobs <- i + 1
	}
	//生产者 放入完毕后要自己关闭
	close(jobs)

	// 扇出 的通道在这里处理 直接这样会阻塞死锁
	//for v := range result {
	//	log.Println("扇出的汇总内容：", v)
	//}

	//这个不灵活
	for i := 0; i < numJobs; i++ {
		log.Println("扇出的汇总内容：", <-result)

	}
}
