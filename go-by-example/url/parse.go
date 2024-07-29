package main

import (
	"log"
	"net"
	"net/url"
	"strings"
)

func main() {
	scheme := "http://username:pwd11@baidu.com:1123/path?k=v#f"
	u, err := url.Parse(scheme)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(u.Scheme, u.User, u.Host, u.Path, u.Fragment)
	log.Println(u.Opaque)
	log.Println(u.RawQuery)
	log.Println(u.RawPath)
	log.Println(u.User.Username())
	log.Println(u.User.Password())
	log.Println(u.User.String())

	log.Println(net.SplitHostPort(u.Host))

	log.Println(u.Path)
	log.Println(u.RawQuery)
	log.Println(u.Parse("k=v"))

	v, err := (url.ParseQuery(u.RawQuery))
	if err != nil {
		log.Println(err.Error())
	}
	for k, v := range v {
		log.Printf("%s,%q", k, strings.Join(v, ""))
	}
}
