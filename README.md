# go-naive-proxy
### Description
This project is a naive Proxy, which can be used as reverse proxy for SSL authentication .<br/>
The server redirect any request to the target address configured.
As for now the target can only use HTTP scheme.
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


