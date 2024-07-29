package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func test() {

	data := "go-by-example"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
func main() {
	//base64Str := base64.StdEncoding.EncodeToString([]byte("go-by-example"))
	//log.Println(base64Str)
	//
	//dst, err := base64.StdEncoding.DecodeString(base64Str)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//log.Println(string(dst))
	//
	//log.Println("-------")
	//base64Str = base64.RawURLEncoding.EncodeToString([]byte("go-by-example"))
	//log.Println(base64Str)
	//
	//dst, err = base64.RawURLEncoding.DecodeString(base64Str)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//log.Println(string(dst))

	log.Println("-------")
	base64Str := base64.URLEncoding.EncodeToString([]byte("go-by-example"))
	log.Println(base64Str)

	dst, err := base64.URLEncoding.DecodeString(base64Str)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(dst))
	log.Println("+++++")

	test()
}
