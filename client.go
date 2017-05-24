package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

var endpoint = []byte("http://api.coindesk.com/v1/bpi/currentprice.json")
var httpclient *http.Client
var proxypath string

func init() {
	flag.StringVar(&proxypath, "socks", "", "socks proxy to use (format: socks5://localhost:1080)")
}

// sample JSON response from API:
// {"time":{"updated":"May 23, 2017 23:30:00 UTC","updatedISO":"2017-05-23T23:30:00+00:00","updateduk":"May 24, 2017 at 00:30 BST"},"disclaimer":"This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org","chartName":"Bitcoin","bpi":{"USD":{"code":"USD","symbol":"&#36;","rate":"2,290.6025","description":"United States Dollar","rate_float":2290.6025},"GBP":{"code":"GBP","symbol":"&pound;","rate":"1,766.9295","description":"British Pound Sterling","rate_float":1766.9295},"EUR":{"code":"EUR","symbol":"&euro;","rate":"2,047.4436","description":"Euro","rate_float":2047.4436}}}

type ResponseTime struct {
	Updated    string    `json:"updated"`
	UpdatedISO time.Time `json:"updatedISO"`
	UpdatedUK  string    `json:"updateduk"`
}

// Bitcoin Price Index
type BPI struct {
	Code        string  `json:"code"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"ratefloat"`
}

type Response struct {
	Time       ResponseTime   `json:"time"`
	Disclaimer string         `json:"disclaimer"`
	ChartName  string         `json:"chartName"`
	BPI        map[string]BPI `json:"bpi"`
	Error      string         `json:",omitempty"`
}

func (r Response) GetTime() string {
	return r.Time.UpdatedISO.Format("Mon Jan 2")

}
func gethttpclient() error {
	client := http.DefaultClient
	if proxypath != "" {
		proxyurl, err := url.Parse(proxypath)
		if err != nil {
			return err
		}
		dialer, err := proxy.FromURL(proxyurl, proxy.Direct)
		if err != nil {
			return err
		}
		transport := &http.Transport{Dial: dialer.Dial}
		client.Transport = transport
	}
	httpclient = client
	return nil
}
func Get() Response {
	resp, err := httpclient.Get(string(endpoint))
	if err != nil {
		return Response{Error: err.Error()}
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{Error: err.Error()}
	}
	var response Response
	err = json.Unmarshal(b, &response)
	if err != nil {
		return Response{Error: err.Error()}
	}
	return response
}
