package main

import (
	"encoding/base64"
	"log"
)

func main() {
	buffer := make([]byte, 1024)
	base64.URLEncoding.Encode(buffer, []byte("go-by-example"))
	log.Println(string(buffer))
}
