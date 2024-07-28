package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//scanner.Buffer()
	//bufio.MaxScanTokenSize

	// ctrl +d 可以 代表EOF
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			//return
			break
		}
		log.Println(strings.ToUpper(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	} else {

		log.Println("没有错误")
	}
	fmt.Println("end")
}
