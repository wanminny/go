package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"
)

func main() {

	cmd := exec.Command("grep", "hello")
	log.Println(cmd)
	log.Println(cmd.String())

	//rlt, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//log.Println("----", string(rlt))

	stdW, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err.Error())
	}
	stdO, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err.Error())
	}
	stdE, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err.Error())
	}

	// 要显示指定开始；前面都是设置 一下输入输出；
	cmd.Start()
	stdW.Write([]byte("hello  3333"))
	stdW.Close() // 必须要关闭 否则一直阻塞

	buffer := make([]byte, 1024)
	stdO.Read(buffer)

	c, err := io.ReadAll(stdE)
	if err != nil {
		log.Println(err.Error())
	}

	// 或者下面方法读取也可以！
	//buffer, err := io.ReadAll(stdO)
	//if err != nil {
	//	log.Println(err.Error())
	//}

	cmd.Wait()
	fmt.Println("----", string(buffer), "===>", string(c))
	time.Sleep(time.Second * 10)
}
