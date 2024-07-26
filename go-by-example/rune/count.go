package main

import (
	"log"
	"unicode/utf8"
)

func count() {
	var s = "怎么才可以赚钱；"
	log.Println(len(s))
	log.Println(utf8.RuneCountInString(s))
	r, size := utf8.DecodeRuneInString("么")
	log.Printf("%c %d", r, size)
	log.Println("==========>")
	for i := 0; i < len(s); {
		rValue, Width := utf8.DecodeRuneInString(s[i:])
		log.Println(rValue, Width)
		log.Printf("%#U %d", rValue, Width)
		i += Width

	}
	log.Println("-----------")

	for i := 0; i < len(s); {
		rValue, Width := utf8.DecodeRuneInString(s[i:])
		//log.Println(rValue, Width)
		log.Printf("%#U %d", rValue, Width)
		i += Width
	}
	for k, v := range s {
		log.Printf("%d %q", k, v)

	}
	log.Println("********")

	for len(s) > 0 {
		r1, w1 := utf8.DecodeRuneInString(s)
		s = s[w1:]
		log.Printf("%q has %d bytes", r1, w1)
		judge(r1)
	}
}

func judge(r rune) {
	if r == '才' {
		log.Printf("has %q\n", r)
	} else {
		log.Println("no ...")
	}
}
func main() {
	count()
}
