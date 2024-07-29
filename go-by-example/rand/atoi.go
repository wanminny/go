package main

import (
	"log"
	"strconv"
)

func main() {
	log.Println(strconv.ParseFloat("1.234", 64))

	log.Println(strconv.ParseInt("22", 0, 64))

	log.Println(strconv.ParseBool("1"))
	log.Println(strconv.ParseBool("0"))
	log.Println(strconv.ParseBool("a"))

	log.Println(strconv.ParseUint("0xFfee", 0, 64))
	log.Println(strconv.ParseUint("0o18", 0, 64))
	log.Println(strconv.ParseUint("0b11", 0, 64))
	log.Println(strconv.ParseUint("11", 0, 64))
	log.Println(strconv.ParseUint("011", 0, 64))
	log.Println(strconv.ParseUint("0o11", 0, 64))

	log.Println(strconv.Atoi("111"))
	log.Println(strconv.Atoi("0x345"))
	log.Println(strconv.ParseUint("0x345", 0, 64))

}
