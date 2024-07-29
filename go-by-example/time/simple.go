package main

import (
	"log"
	"time"
)

func main() {

	log.Println(time.Now())
	log.Println(time.Now().Second())

	log.Println(time.Now().Year())
	log.Println(time.Now().Month())
	log.Println(time.Now().Day())
	log.Println(time.Now().Hour())
	log.Println(time.Now().Minute())
	log.Println(time.Now().Second())
	log.Println(time.Now().Nanosecond())
	log.Println(time.Now().Location(), time.Local)
	log.Println(time.Now().Weekday())
	log.Println(time.Now())

	//time.Now().Equal()
	//time.Date()
	log.Println("-------------")
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	log.Println(then)
	log.Println(then.Before(time.Now()))
	log.Println(then.After(time.Now()))
	log.Println(then.Equal(time.Now()))

	diff := time.Now().Sub(then)
	log.Println(diff)

	log.Println(diff.Hours())
	log.Println(diff.Minutes())
	log.Println(diff.Seconds())
	log.Println(diff.Nanoseconds())

	log.Println(then.Add(diff))
	log.Println(then.Add(-diff))
}
