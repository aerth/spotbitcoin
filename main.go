package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	addr := "0.0.0.0:8080"
	if a := os.Getenv("ADDR"); a != "" {
		println("using address:", a)
		addr = a
	}
	if p := os.Getenv("PORT"); p != "" {
		println("using port:", p)
		addr = "0.0.0.0:" + p
	}
	if flagport != "" {
		println("using address:", flagport)
		addr = "0.0.0.0:" + flagport
	}
	gethttpclient()
	s := NewSystem()
	h := http.DefaultServeMux
	h.Handle("/", http.HandlerFunc(home))
	h.Handle("/current.png", s)
	println("listening", addr)
	err := http.ListenAndServe(addr, h)
	if err != nil {
		println(err.Error())
	}
	os.Exit(111)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html>
<html>
<img style="max-width: 100%; height: auto;" src="/current.png">
<br>Powered by <a href="https://coindesk.com">CoinDesk<a> API | 
<a href="https://github.com/aerth/spotbitcoin/">source code (MIT)</a> 
by aerth
</html>
`))
}
