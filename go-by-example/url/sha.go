package main

import (
	"crypto/sha256"
	"log"
)

func main() {
	h := sha256.New()

	// 写入要处理的字符串
	h.Write([]byte("go by example"))
	log.Printf("%x\n", h.Sum(nil))
}
