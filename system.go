package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

var CacheTime = time.Second * 60

type System struct {
	mu    sync.Mutex
	last  time.Time
	cache Response
}

func NewSystem() *System {
	s := new(System)
	s.last = time.Now()
	s.cache = Get()
	return s
}

func (s *System) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}
	res := s.cache
	s.mu.Lock() // one at a time (still quick)
	if time.Now().Sub(s.last) > CacheTime {
		fmt.Println("\nGrabbing from API\n")
		res = Get()
		if res.Error != "" {
			println(time.Now().String(), res.Error)

		}
		s.last = time.Now()
		s.cache = res

	} else {
		fmt.Println("Using Cached response from:", s.last)
	}
	s.mu.Unlock()

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
