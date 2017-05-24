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
	h.Handle("/current.png", s)
	h.Handle("/", http.HandlerFunc(home))
	http.ListenAndServe(addr, h)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/current.png"))
}
