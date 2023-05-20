package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is a FastCGI example server.\n"))
}

func main() {
	e := "/tmp/fcgi-example.sock"

	os.Remove(e)

	fmt.Println("Starting server on", e, "...")

	l, err := net.Listen("unix", e)

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
