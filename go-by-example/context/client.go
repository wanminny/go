package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("start process req")
	defer fmt.Println("end process req HelloHandler ")

	ctx := req.Context()

	select {
	case <-ctx.Done():
		fmt.Println("cancel .....done!!!")
		err := ctx.Err()
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case <-time.After(time.Second * 4):
		fmt.Println("after 4 s")
		//w.Write([]byte("hello"))
		fmt.Fprintf(w, "%s", "hello....")

	}
	// 是否必须
	//return
}
func main() {

	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8081", nil)

}
