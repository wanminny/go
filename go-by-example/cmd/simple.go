package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	//flag.Parse()
	fmt.Println(os.Args, os.Args[0])

	fmt.Println("---------")

	fmt.Println(flag.Args())

	port := flag.Int("port", 8080, "port")
	url := flag.String("url", "baidu.com", "url")
	log.Println(flag.Args())
	all := flag.Bool("all", false, "-A or all")
	var bare string
	flag.StringVar(&bare, "b", "bare", "--bare")
	flag.Parse()
	log.Println(*port, *url, *all, bare)

	// 表示参数 需要在最后否则会把后面的flag也做完parameter 
	fmt.Println(flag.Args())

	// 多少个flag参数
	fmt.Println(flag.NFlag())

}
