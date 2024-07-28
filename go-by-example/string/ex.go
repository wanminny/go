package main

import (
	"fmt"
	"os"
	"strings"
)

var p = fmt.Println

func test() {
	p("contains:", strings.Contains("go by example", "by"))
	p("count: ", strings.Count("go-by-example", "e"))
	p("count: ", strings.Count("go-by-example", "-"))
	p("count: ", strings.Count("go-by-example", "-by"))
	p("hasPrefix", strings.HasPrefix("go-by-基本库", "by"))
	p("hasPrefix", strings.HasPrefix("go-by-基本库", "go"))
	p("hasSuffix", strings.HasSuffix("go-by-基本数据结构", "数据结构"))
	p("join", strings.Join([]string{"go", "by", "基本算法"}, "*"))
	p("Index", strings.Index("go by example", "by"))
	p("split", strings.Split("go by example", " "))
	fmt.Println(strings.Repeat("go", 3))
	fmt.Println(strings.Replace("go by go example", "go", "golang", 1))
	fmt.Println(strings.Replace("go by go example go ", "go", "golang", 2))
	fmt.Println(strings.Replace("go by go example go ", "go", "golang", -11))
	fmt.Println(strings.ReplaceAll("go by go example go ", "go", "golang"))
	p("gobyexample"[2])
}

type point struct {
	x, y int
}

func test1() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Println("----------")

	fmt.Fprintf(os.Stderr, "%s\n", "go by ex error")
	tmp := fmt.Sprintf("%s, \n", "go by ,,,")
	fmt.Println(tmp)

	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", false)

	fmt.Printf("%d,%b,%x\n", 15, 15, 15)
	fmt.Printf("%d\n", 'A')

	fmt.Printf("%.3f\n", 88.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"go by ex\"")
	fmt.Printf("%q\n", "go by ex")
	fmt.Printf("%q\n", "\"go by ex\"")

	fmt.Printf("%x\n", "go by")

	fmt.Println("===========================")
	fmt.Printf("|%6d|%d|%-6d|\n", 23, 45, 67)
	fmt.Printf("|%6s|%s|%-6s|\n", "23", "45", "67")

	fmt.Printf("|%6.2f|%6.2f|%7.3f|\n", 12.2, 12.2, 33.56)
	fmt.Printf("|%-6.2f|%-6.2f|%-7.3f|\n", 12.2, 12.2, 33.56)

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	fmt.Printf("width5: |%6s|%6s|\n", "foo", "b")

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width2: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("%x", "ab")
	fmt.Printf("str3: %x\n", "hex this")

}
func main() {
	//test()

	test1()
}
