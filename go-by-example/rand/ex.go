package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println(rand.Intn(100), ",")
	log.Println(rand.Intn(100))

	log.Println(rand.Intn(6), ",")
	log.Println(rand.Intn(6))

	log.Println(rand.Float64(), ",")
	log.Println(rand.Float64(), ",")
	log.Println(rand.Float64(), ",")

	log.Println(rand.Float64()*5+5, ",")

	log.Println("-------")
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	fmt.Println(r.Intn(10), ",")
	fmt.Println(r.Intn(10))
	log.Println("-------")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(3))
	fmt.Println(r1.Intn(3))
	log.Println("---++++----")

	s2 := rand.NewSource(44)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(44)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))

}
