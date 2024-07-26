package main

import (
	"fmt"
	"math"
	"time"
)

const AA = 1

var packageFloat float64

func value() {
	fmt.Println("22 +33= ", 22+33)
	fmt.Println("go" + "by" + "example!")
	fmt.Println("7.0/3.0= ", 7.0/3.0)
	fmt.Println(true && true)
	fmt.Println("false && true= ", false && true)
	fmt.Println("false || true=", false || true)
	fmt.Println("!false=", !false)

}

func variable() {
	var a = 1
	b := 1
	var s = "go by example"
	d := true
	var e int
	var a1, b1 = 22, 33
	fmt.Println(a, b, s, d, e, AA, packageFloat, a1, b1)
	fmt.Printf("%.2f\n", packageFloat)

}

const TOKEN = "constant"

func constVariable() {
	fmt.Println(TOKEN)
	const n = 500000
	const d = 3e10 / n
	fmt.Println(d)

	fmt.Printf("%f\n", d)
	fmt.Println(math.Sin(d))
}
func forI() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		} else {
			fmt.Println(i)
		} // 0 1
	}
	for {
		fmt.Println("for loop")
		break // for loop
	}
	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i) // 1 3
	}
}
func ifElseIt() {
	if num := 4; num < 2 {
		fmt.Println("num < 2")
	} else if num > 8 {
		fmt.Println("num gt 8")
	} else {
		fmt.Println("兜底。。。")
	}
}

func swCase() {
	timeChoice := time.Now().Weekday()
	switch timeChoice {
	case time.Thursday:
		fmt.Println(time.Thursday)
	case time.Friday:
		fmt.Println(time.Friday)
	case time.Saturday:
		fmt.Println(time.Saturday)
	}
	timeHour := time.Now().Hour()
	switch timeHour {
	case 16:
		fmt.Println(timeHour)
		fallthrough
	case 17:
		fmt.Println("0000-", timeHour)
	}

	whatAmi := func(w any) {
		//typ := w.(type)
		switch typ := w.(type) {
		case string:
			fmt.Println(typ)
		case bool:
			fmt.Println(typ)
		default:
			fmt.Printf("not support %T\n", typ)
		}
	}
	whatAmi("go by example")
	whatAmi(false)
	whatAmi(1)
}
func main() {
	fmt.Println("hello go by example！")
	value()
	fmt.Println("=========")
	variable()
	fmt.Println("=========")
	constVariable()
	fmt.Println("=========")
	forI()
	fmt.Println("---------->")
	ifElseIt()
	fmt.Println("++++++++")
	swCase()

}
