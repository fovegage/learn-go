package main

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
)

func c15ProxyParse() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8087")
	}
	println(proxy)
}

func c15ProxyT1() {
	ctx := context.TODO()
	ipEchoServer := "84.17.45.179"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ipEchoServer, nil)
	if err != nil {

	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	if data, err := io.ReadAll(resp.Body); err != nil {

	} else {
		println(net.ParseIP(string(data)))
	}
}

func main() {
	c15ProxyT1()
}
