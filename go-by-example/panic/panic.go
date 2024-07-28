package main

import (
	"fmt"
	"os"
)

func myPanic() {
	panic("go by example...")
}
func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprintf(os.Stderr, "=====>%q", err)
		}
	}()

	myPanic()
	fmt.Println("after panic ..")
}
