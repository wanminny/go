package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("====>", err)
		}
	}()
	panic("end")

	defer func() {
		log.Println(2222)
	}()

	defer fmt.Println(11111)

	f, err := os.Create("/tmp/aa.txt")
	if err != nil {
		panic(err.Error())
	} else {
		f.Close()
	}

}
