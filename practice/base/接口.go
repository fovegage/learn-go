package main

import (
	"../pkg/mock"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func downloader(r Retriever) string {
	// 在调用get之前可以做一些事情
	// 业务代码只需要实现get方法即可
	fmt.Println("执行")
	return r.Get("https://www.gaozhe.net")
}

// 测试代码
type middleWare struct {
	r mock.Retriever
}

func (m middleWare) Get(url string) {
	m.r.Get(url)
}

func NewMiddleWare(r string) {
	m := middleWare{}
	fmt.Println(m.r.Get("111"))
}

func main() {
	// 有点向上转型的意思
	var r mock.Retriever
	// 在 downloader 里面的东西必须实现 get方法
	fmt.Println(downloader(r))

	// 测试2
	//r := mock.Retriever{"hello duck"}
	//fmt.Println(r)
	//fmt.Println(r.Get("2"))

	//NewMiddleWare("hello")
}
