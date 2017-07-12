package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"net/http/httputil"
	"net/url"
)

type httpResponse func (rw http.ResponseWriter, req *http.Request)
func handleHttp(reverseProxy *httputil.ReverseProxy) httpResponse {
	return func (rw http.ResponseWriter, req *http.Request) {
		fmt.Println("got request")
		reverseProxy.ServeHTTP(rw, req)
	}
}

func getProxyHandler(targetUrl string) (*httputil.ReverseProxy, error){
	if len(targetUrl) == 0 {
		targetUrl = "http://localhost:8081/"
	}
	fmt.Println("target url:", targetUrl)
	var err error
	var target *url.URL
	target, err = url.Parse(targetUrl)
	if err != nil {
		return nil, err
	}
	origHost := target.Host
	d := func(req *http.Request) {
		req.URL.Host = origHost
		req.URL.Scheme = "http"
	}

	p := &httputil.ReverseProxy{Director: d,}
	return p, nil
}

func main() {
	handler, err := getProxyHandler("")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handleHttp(handler))
	go func() {
		err := http.ListenAndServe(":8099", nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("Connected http")
	}()

	absPathCert, _ := filepath.Abs("cert.pem")
	absPathKey, _ := filepath.Abs("key.pem")
	err = http.ListenAndServeTLS(":444", absPathCert, absPathKey, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected https")
}
