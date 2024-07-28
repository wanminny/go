package main

import (
	"fmt"
	"os"
)

func createFile(name string) (*os.File, error) {
	f, err := os.Create(name)
	if err != nil {
		panic(err.Error())
	}
	return f, nil
}

func writeFile(file *os.File, b []byte) {
	_, err := file.Write(b)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func closeFile(file *os.File) {
	fmt.Println("close file ....")
	err := file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	f, err := createFile("/tmp/defer.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	writeFile(f, []byte("go-by-example"))
	closeFile(f)

}
