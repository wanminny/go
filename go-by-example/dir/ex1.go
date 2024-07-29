package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func checkerr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func main() {
	fInfo, err := os.Stat("subdir")
	//checkerr(err)
	if fInfo == nil {
		err = os.Mkdir("subdir", 0755)
		checkerr(err)
	}
	//if !fInfo.IsDir() {
	//}
	defer func() {

		err := os.RemoveAll("subdir")
		checkerr(err)
	}()
	createEmptyFile := func(name string) {
		err = os.WriteFile(name, []byte(""), 0755)
		checkerr(err)
		//f, err := os.Create(name)
		//checkerr(err)
		//f.Write([]byte(""))
	}

	createEmptyFile("subdir/file1")
	createEmptyFile("subdir/file2")

	err = os.MkdirAll("subdir/parent/child", 0755)
	checkerr(err)
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")
	dirs, err := os.ReadDir("subdir/parent")
	checkerr(err)
	for k, v := range dirs {
		info, _ := v.Info()
		log.Println(k, v.Name(), v.IsDir(), v.Type(), info.ModTime())
	}
	err = os.Chdir("subdir/parent/child")
	checkerr(err)

	dirs, err = os.ReadDir(".")
	checkerr(err)
	for k, v := range dirs {
		log.Println(k, v)
	}
	os.Chdir("../../../")

	dirs, err = os.ReadDir(".")
	checkerr(err)
	for k, v := range dirs {
		log.Println(k, v)
	}
	log.Println("++++++")
	err = filepath.Walk("subdir", visit)
	if err != nil {
		log.Println(err.Error())
	}

}

func visit(path string, info fs.FileInfo, err error) error {
	//err = fmt.Errorf("%s", "file err ！")
	if err != nil {
		return err
	}
	fmt.Println(path, info.IsDir(), info.Name(), info.Size())
	return nil
}
