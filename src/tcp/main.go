package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is a FastCGI example server.\n"))
}

func main() {
	e := "localhost:1111"

	fmt.Println("Starting server on", e, "...")

	l, err := net.Listen("tcp", e)

	if err != nil {
		panic(err)
	}

	defer l.Close()

	b := new(FastCGIServer)

	err = fcgi.Serve(l, b)

	if err != nil {
		panic(err)
	}
}
