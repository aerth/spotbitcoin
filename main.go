package main

import (
	"flag"
	"net/http"
)

func main() {
	flag.Parse()
	addr := "0.0.0.0:80"
	gethttpclient()
	s := new(System)
	h := http.DefaultServeMux
	h.Handle("/", s)
	http.ListenAndServe(addr, h)

}
