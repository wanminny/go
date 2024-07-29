package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//flag.Parse()
	fooSet := flag.NewFlagSet("sub", flag.ExitOnError)
	fooStr := fooSet.String("foo", "foo abc", "foo 参数")
	fooPort := fooSet.Int("port", 8080, " sub命令的port")

	barSet := flag.NewFlagSet("clone", flag.ExitOnError)
	barStr := barSet.String("bar", "bar abc", "bar 参数")
	barPort := barSet.Int("port", 9090, "clone 命令的端口号")

	// 这个有里才可以解析 父命令行本身！！
	flag.Parse()

	if len(os.Args) < 2 {
		log.Println(" sub or clone 子命令，应该是 xxx  sub or xx clone ")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "sub":
		fooSet.Parse(os.Args[2:])
		log.Println(*fooStr)
		log.Println(*fooPort)
		log.Println(os.Args)
		log.Println(flag.NFlag())
		log.Println(flag.Args())
		log.Println("+++++++++++")
		log.Println(fooSet.NFlag())
		log.Println(fooSet.Args())

	case "clone":
		barSet.Parse(os.Args[2:])
		log.Println(*barStr)
		log.Println(*barPort)
		log.Println(os.Args)
		log.Println(flag.NFlag())
		log.Println(flag.Args())
		log.Println("+++++++++++")
		log.Println(barSet.NFlag())
		log.Println(barSet.Args())

	default:
		fmt.Println("输入子命令错误 available subcmd: sub/clone")
		os.Exit(1)

	}
}
