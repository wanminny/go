package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func main() {
	data, err := os.ReadFile("/tmp/defer.txt")
	checkerr(err)

	log.Println(string(data))

	//os.OpenFile()
	f, err := os.Open("/tmp/defer.txt")
	checkerr(err)
	defer f.Close()

	buffer := make([]byte, 512)
	n, err := f.Read(buffer)
	checkerr(err)
	log.Printf("%d bytes, 内容是:%s", n, buffer[0:n+1])
	fmt.Println("------分隔符-------")
	off1, err := f.Seek(6, 0)
	checkerr(err)
	buf1 := make([]byte, 2)
	n1, err := io.ReadAtLeast(f, buf1, 2)
	checkerr(err)

	fmt.Printf("%d bytes from %d: %s\n", n1, off1, string(buf1))

	n2, err := f.Seek(0, 0)
	checkerr(err)

	r := bufio.NewReader(f)
	con1, err := r.Peek(5)
	checkerr(err)
	log.Printf("n2=%d,%s\n", n2, con1)
}
