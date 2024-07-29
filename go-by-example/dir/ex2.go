package main

import (
	"log"
	"os"
	"path"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	f, err := os.CreateTemp("", "sample")
	checkerr(err)
	println(f.Name())

	defer os.Remove(f.Name())

	f.Write([]byte("go by example"))

	tmpDir, err := os.MkdirTemp("", "sample-temp")
	checkerr(err)
	defer os.RemoveAll(tmpDir)
	log.Println(tmpDir)

	fN := path.Join(tmpDir, "file1")
	os.WriteFile(fN, []byte("ccc dd ee"), 0755)

}
