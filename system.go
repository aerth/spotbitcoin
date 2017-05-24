package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type System struct {
	mu sync.Mutex
}

func (s *System) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
	}
	res := Get()
	if res.Error != "" {
		println(time.Now().String(), res.Error)
	}
	fmt.Println(res)
	img, err := drawpng(fmt.Sprintf("%s: 1 Bitcoin (BTC) is currently worth USD %s", res.GetTime(), "$"+strings.Split(res.BPI["USD"].Rate, ".")[0]))
	if err != nil {
		println(err.Error())
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Content-Type", "image/png")
	w.Header().Add("X-Powered-By", "CoinDesk - https://coindesk.com")
	WritePNG(img, w)
	w.Write([]byte(`Powered by <a href="https://coindesk.com/price/">CoinDesk</a>`))
}
