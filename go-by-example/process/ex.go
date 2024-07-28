package main

import (
	"log"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("bash", "-c", "ls -lah")

	//rlt, err := dateCmd.CombinedOutput()
	rlt, err := dateCmd.Output()
	if err != nil {
		log.Println(err.Error(), "-->", string(rlt))
		return
		//panic("执行错误！")
	}

	//err := dateCmd.Run()
	//if err != nil {
	//	log.Println(err.Error())
	//	panic("执行错误！")
	//}
	//rlt, _ := dateCmd.Output()
	//rlt, _ = dateCmd.CombinedOutput()
	log.Println(dateCmd.String())
	log.Println(string(rlt))
}
