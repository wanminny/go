package main

import (
	"errors"
	"fmt"
	"log"
)

type SEX uint8

const (
	Male SEX = iota
	FeMale
	Secret
)

var ERR_BY_ZERO = errors.New("err by zero ！")
var NOT_FOUND = errors.New("not found")

func divided(d int) error {
	if d == 0 {
		return ERR_BY_ZERO
	} else {
		return nil
	}
}
func findRecord(sex SEX) (err error) {
	if sex == Secret {
		return NOT_FOUND
	} else {
		return nil
	}
}

func warp(err error) error {
	return fmt.Errorf("warped %s", "--->")
}

// 只有这个才是包装过的使用%w 专门为 error而做的
func warp1(err error) error {
	return fmt.Errorf("warped %s  %w", "--->", err)
}
func main() {
	err := divided(0)
	if err == ERR_BY_ZERO {
		log.Println("除数为0了！")
	}
	log.Println(warp(err))
	err1 := warp(err)
	err2 := warp1(err)

	if errors.Is(err1, err) {
		log.Println("err1 is  err ")
	}
	if errors.Is(err2, err) {
		log.Println("err2 is  err ")
	}

	log.Println(errors.Unwrap(err2))
	log.Println(errors.Unwrap(err1))
	log.Println("*********")

	//var asErr error
	if errors.As(err2, &err1) {
		log.Println(err1.Error())
		log.Println("as true")
	} else {
		log.Println("as false")
	}

	log.Println(err)
}
