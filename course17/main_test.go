package main

import (
	"bytes"
	"fmt"
	"net/http"
	"study/course17/middleware"
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

func TestChain(t *testing.T) {
	// https://zhuanlan.zhihu.com/p/507085115
	cc := &middleware.MyContext{}
	m1 := func(ctx *middleware.MyContext) {
		//fmt.Println("im m1 begin")
		//ctx.Abort()
		//ctx.Next()
		//fmt.Println("im m1 end")
		fmt.Println("start")
		req, _ := http.NewRequest("https://www.baidu.com", http.MethodPost, bytes.NewBufferString("s"))
		cc.Req = req
		ctx.Next()
		fmt.Println("end")

	}
	m2 := func(ctx *middleware.MyContext) {
		fmt.Println("im m2 begin")
		ctx.Next()
		fmt.Println("im m2 end")
	}
	m3 := func(ctx *middleware.MyContext) {
		fmt.Println("im m3 begin")
		//ctx.Abort()
		ctx.Next()
		fmt.Println("im m3 end")
	}

	//
	mainFunc := func(ctx *middleware.MyContext) {
		fmt.Println(cc.Req)
		fmt.Println("im get function")
	}
	ctx := &middleware.MyContext{}
	ctx.Use(m1)
	ctx.Use(m2)
	ctx.Use(m3)
	ctx.Get("hahahah", mainFunc)
	// 后续怎么处理
	ctx.Run()
}
