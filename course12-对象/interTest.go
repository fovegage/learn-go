package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// NewReader 实现了Read方法
	bytes.NewReader([]byte("hello word"))

	// 规定传递io.Reader()
	goquery.NewDocumentFromReader(bytes.NewReader([]byte("ss")))
	// io.Reader()
}
