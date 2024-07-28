package main

import (
	"log"
	"path/filepath"
	"strings"
)

func main() {

	f := filepath.Join("dir1", "dir2", "file1")
	log.Println(f)
	log.Println(filepath.Base("/a/b/c.txt"), filepath.Dir("/a/b/c.txt"), filepath.Ext("/a/b/c.txt"))

	log.Println("-----")
	log.Println(filepath.Join("/dir//", "file2"))
	log.Println(filepath.Join("/dir/../dir2/", "file3"))

	log.Println(filepath.Dir(f), filepath.Base(f), filepath.Ext(f))

	log.Println(filepath.IsAbs("/a/b/c"), filepath.IsAbs("a/b/c"), filepath.ListSeparator)

	file := "config.json"
	ext := filepath.Ext(file)

	log.Println("==========", ext)

	log.Println(strings.TrimSuffix(file, ext))

	// 理解:
	// 从 `a/b` 进入 `/a/b/c/d.txt`，只需要进入 `c/d.txt` 即可
	s, err := filepath.Rel("/a/b", "/a/b/c/d.txt")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(s)

	// 理解:
	// 从 `a/b` 进入 `/a/f/c/d.txt`，只需要进入 `../f/c/d.txt` 即可
	s, err = filepath.Rel("/a/b", "/a/f/c/d.txt")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(s)

	// 要求 targpath 和 basepath 必须“都是相对路径”或“都是绝对路径”。
	// 要么全部是相对路径要么全部是绝对路径
	s, err = filepath.Rel("d/b", "/a/f/c/d.txt")
	if err != nil {
		log.Println(err.Error()) // Rel: can't make /a/f/c/d.txt relative to d/b
	}
	log.Println(s)

}
