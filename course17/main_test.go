package main

import (
	"study/course17/option"
	"testing"
)

func TestNormal(t *testing.T) {
	// https://www.liwenzhou.com/posts/Go/functional-options-pattern/
	opt := option.NewOptions(option.WithAge(29))
	println(opt.Age, opt.Name)

	age := option.WithAge(29)
	age(&option.Options{
		Name: "gage",
	})
}

func TestInterface(t *testing.T) {
	// 这里即golang版本构造函数的实现
	// https://www.liwenzhou.com/posts/Go/functional-options-pattern/
	server := option.BuildServer(option.WithHost("127.0.0.1"))
	println(server.Host, server.Port)
}
