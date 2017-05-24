package main

import (
	"flag"
	"net/http"
)

func main() {
	flag.Parse()
	addr := "127.0.0.1:8080"
	gethttpclient()
	s := new(System)
	h := http.DefaultServeMux
	h.Handle("/", s)
	http.ListenAndServe(addr, h)

}
