package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	addr := "0.0.0.0:80"
	if a := os.Getenv("ADDR"); a != "" {
		addr = a
	}
	gethttpclient()
	s := new(System)
	h := http.DefaultServeMux
	h.Handle("/current.png", s)
	h.Handle("/", http.HandlerFunc(home))
	println("listening", addr)
	err := http.ListenAndServe(addr, h)
	if err != nil {

		println(err.Error())
	}
	os.Exit(111)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/current.png"))
}
