package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
)

func main() {
	// CGO_ENABLED=1 go run exp1.go
	easy := curl.EasyInit()
	defer easy.Cleanup()

	// 使用隧道代理？还是随机取一个用？

	// export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890

	//easy.Setopt(curl.PROXY_SOCKS5, "socks5://127.0.0.1:7890")

	easy.Setopt(curl.OPT_URL, "https://httpbin.org/ip")

	// easy.Setopt(curl.OPT_URL, "https://www.google111111.com/")

	// make a callback function
	fooTest := func(buf []byte, userdata interface{}) bool {
		println("DEBUG: size=>", len(buf))
		println("DEBUG: content=>", string(buf))
		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
