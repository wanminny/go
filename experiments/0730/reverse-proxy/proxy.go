package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// run nginx on local by docker, and listen 80 port
	proxy, err := NewProxy("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", ProxyRequestHandler(proxy))

	fmt.Println("Starting proxy server :8090; please modify to proxy you want!")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	// Request
	proxy := httputil.NewSingleHostReverseProxy(url)
	originalDirector := proxy.Director
	proxy.Director = func(r *http.Request) {
		originalDirector(r)
		modifyRequest(r)
	}

	// Response 这里是要拦截的！！如果你要自定义！！
	//proxy.ModifyResponse = func(r *http.Response) error {
	//	return errors.New("just a error")
	//}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println("get err:", err)
	}

	return proxy, nil
}

func modifyRequest(req *http.Request) {
	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		proxy.ServeHTTP(w, req)
	}
}
