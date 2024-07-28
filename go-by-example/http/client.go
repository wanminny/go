package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://gobyexample-cn.github.io/")
	if err != nil {
		log.Println(err.Error())
	}

	defer resp.Body.Close()
	fmt.Println("status:", resp.Status)

	//s := bufio.NewScanner(resp.Body)
	//for s.Scan() {
	//	fmt.Println("txt:", s.Text())
	//}
	//
	//if err = s.Err(); err != nil {
	//	log.Println("err: ", err.Error())
	//}

	c, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(c))
}
