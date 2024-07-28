package main

import (
	"bufio"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	err := os.WriteFile("/tmp/dat1", []byte("hello11,"), 0o644)
	check(err)

	f, err := os.Create("/tmp/ccc.txt")
	check(err)

	defer f.Close()
	d := []byte{68, 69, 70}
	n1, err := f.Write(d)
	check(err)
	log.Printf("write %d %s", n1, string(d))

	n2, err := f.WriteString("go by example\n")
	check(err)
	log.Printf("%d %d", n2, len("go by example\n"))
	f.Sync()

	//针对读写操作都可以seek
	f.Seek(0, 0)

	w := bufio.NewWriter(f)
	n3, err := w.WriteString("数据结构\n")
	check(err)
	log.Println(n3)
	// bufio必须 flush否则没有写入
	w.Flush()

	//r := bufio.NewReader(f)
	//r.Read()

}
