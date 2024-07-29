package main

import (
	"fmt"
	"log"
	"time"
)

func test() {
	// 获取当前时间
	now := time.Now()

	// 获取当前时区信息
	name, offset := now.Zone()
	fmt.Printf("当前时区: %s, UTC 偏移量: %d\n", name, offset)

	// 解析特定时区的时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("加载时区失败:", err)
		return
	}
	t := time.Now().In(loc)
	fmt.Println("中国标准时间:", t)
}
func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))

	t1, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(t1)
	p := fmt.Println
	p(t.Format("3:04PM"))

	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, err := time.Parse(form, "8 41 PM")
	p(t2, err)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)

	test()
}
