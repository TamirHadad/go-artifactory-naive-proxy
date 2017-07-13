package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"net/http/httputil"
	"net/url"
	"github.com/TamirHadad/go-artifactory-naive-proxy/server/cert"
	"os"
)

type httpResponse func(rw http.ResponseWriter, req *http.Request)

func handleHttp(reverseProxy *httputil.ReverseProxy) httpResponse {
	return func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("*********************************************************")
		fmt.Println("Host:    ", req.Host)
		fmt.Println("Method:  ", req.Method)
		fmt.Println("URI:     ", req.RequestURI)
		fmt.Println("Agent:   ", req.UserAgent())
		fmt.Println("*********************************************************")
		reverseProxy.ServeHTTP(rw, req)
	}
}

func getProxyHandler(targetUrl string) (*httputil.ReverseProxy, error) {
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

	p := &httputil.ReverseProxy{Director: d, }
	return p, nil
}

func getProxyTarget() string {
	proxyTarget := ""
	if len(os.Args) > 1 {
		proxyTarget = os.Args[1]
	}
	return proxyTarget
}

func startHTTPListener(handler *httputil.ReverseProxy) {
	go func() {
		// We can use the same handler for both HTTP and HTTPS
		httpMux := http.NewServeMux()
		httpMux.HandleFunc("/", handleHttp(handler))
		port := "8099"
		if httpPort := os.Getenv("PROXY_HTTP_PORT"); httpPort != "" {
			port = httpPort
		}
		err := http.ListenAndServe(":" + port, httpMux)
		if err != nil {
			panic(err)
		}
	}()
}

func prepareHTTPSHandling(handler *httputil.ReverseProxy) (*http.ServeMux, string, string) {
	// We can use the same handler for both HTTP and HTTPS
	httpsMux := http.NewServeMux()
	httpsMux.HandleFunc("/", handleHttp(handler))
	if _, err := os.Stat(cert.CERT_FILE); os.IsNotExist(err) {
		cert.CreateNewCert()
	}
	absPathCert, _ := filepath.Abs(cert.CERT_FILE)
	absPathKey, _ := filepath.Abs(cert.KEY_FILE)
	return httpsMux, absPathCert, absPathKey
}

func main() {
	proxyTarget := getProxyTarget()
	handler, err := getProxyHandler(proxyTarget)
	if err != nil {
		panic(err)
	}

	// Starts a new Go routine
	startHTTPListener(handler)
	httpsMux, absPathCert, absPathKey := prepareHTTPSHandling(handler)
	port := "444"
	if httpPort := os.Getenv("PROXY_HTTPS_PORT"); httpPort != "" {
		port = httpPort
	}
	err = http.ListenAndServeTLS(":" + port, absPathCert, absPathKey, httpsMux)
	if err != nil {
		panic(err)
	}
}
