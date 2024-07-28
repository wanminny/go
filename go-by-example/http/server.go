package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/header", HeadHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func HeadHandler(writer http.ResponseWriter, request *http.Request) {
	for name, header := range request.Header {
		for k, v := range header {
			log.Printf("%s,%v,%v", name, k, v)
			fmt.Fprintf(writer, "%v,%v\n", name, v)
		}
	}

}

//func HeadHandler(writer http.ResponseWriter, request *http.Request) {
//	header := request.Header
//	for k, v := range header {
//		log.Println(k, v)
//	}
//	rlt, err := json.Marshal(header)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	fmt.Fprint(writer, string(rlt))
//}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("ok"))
}
