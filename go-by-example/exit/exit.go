package main

import (
	"log"
	"os"
)

func main() {

	defer func() {
		log.Println("defer execute.....")
	}()
	return
	os.Exit(3)
}
