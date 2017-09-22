package main

import (
	"net"
	"net/http"
	"time"
)

func main() {
	c := http.DefaultClient
	t := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,
	}
	c.Transport = t

	for i := 0; i < 1000000; i++ {
		c.Get("http://localhost:1323/")
	}
}
