# go-naive-proxy
## Description
This project is a naive Proxy, which can be used as reverse proxy for SSL authentication .<br/>
The server redirect any request to the target address configured.
As for now the target can only use HTTP scheme.

##Building the Executable
The project is written in the [Go programming language](https://golang.org/), so to build and run by yourself, <br/>
you first need to have Go installed and configured on your machine.
## Setup Go
To download and install Go, please refer to the [Go documentation](https://golang.org/doc/install).

Navigate to the directory where you want to create the jfrog-cli-go project, and set the value of the GOPATH environment variable to the full path of this directory.
## Download and Build
```
go get github.com/TamirHadad/go-proxy/server
```
Go will download and build the project on your machine.<br/>
 Once complete, you will find the server executable under your $GOPATH/bin directory.

### SSL Configuration
The server will load naive_proxy_cert.pem naive_proxy_key.pem. <br/>
If the pem files are missing they will be created.<br/>
By default the server will start by using 8099 and 444 port.<br/>
The Ports are configurable using the following environment variables:
 ```
 export PROXY_HTTP_PORT=8080
 ```
 ```
 export PROXY_HTTPS_PORT=443
 ```

## Starting The Server
By default the target of the proxy is `http://localhost:8081/`
```
./server
```
 You can change the default target by running the command with the target server as argument
 ```
 ./server http://127.0.0.1:8082/
 ```


