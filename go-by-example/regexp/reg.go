package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	b, err := regexp.MatchString("p([a-z]+)ch", "pxyzch")
	log.Println(b, err)

	r, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(r.MatchString("peach"))
	log.Println(r.MatchString("pEAch"))

	log.Println("------")
	log.Println(r.FindString("punch peach "))
	log.Println(r.FindStringIndex(" peach puch"))
	log.Println(r.FindStringSubmatch("peach punch "))
	log.Println(r.FindStringSubmatchIndex("peach punch "))

	log.Println(r.FindAllString("peach punch pinch", -1))
	log.Println(r.FindAllString("peach punch pinch", -1))

	log.Println(r.FindAllStringIndex("peach punch pinch", -1))

	log.Println(r.FindAllStringSubmatch("peach punch pinch", 1))

	log.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	log.Println(r.Match([]byte("peach punch pinch")))

	reg := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", reg)

	log.Println(reg.ReplaceAllString("a peach", "<fruits>"))
	log.Println(string(r.ReplaceAllFunc([]byte("a peach!"), func(bytes []byte) []byte {
		return []byte(strings.ToUpper(string(bytes)))
	})))
	log.Println(string(r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)))
}
