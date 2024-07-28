package main

import (
	"log"
	"os"
	"strings"
)

func main() {

	os.Setenv("go", "by  example")
	os.Setenv("orm", "gorm")

	println(os.Getenv("go"), os.Getenv("orm"))

	log.Println("--------------")
	envs := os.Environ()
	for _, v := range envs {
		key := strings.SplitN(v, "=", 2)[0]
		value := strings.SplitN(v, "=", 2)[1]
		log.Println(key, "==>", value)
		//log.Println(k, v)
	}
}
