package main

import (
	"log"
	"strconv"
	"syscall/js"
)

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

var (
	document = js.Global().Get("document")
	numEle   = document.Call("getElementById", "num")
	ansEle   = document.Call("getElementById", "ans")
	btnEle   = js.Global().Get("btn")
)

func fibFunc(this js.Value, args []js.Value) interface{} {
	v := numEle.Get("value")
	if num, err := strconv.Atoi(v.String()); err == nil {
		// 设置html
		ansEle.Set("innerHTML", js.ValueOf(fib(num)))
		// todo:
		log.Printf("test 1111")

	}
	return nil
}

func main() {
	//通过 js.Global().Get("btn") 或 document.Call("getElementById", "num") 两种方式获取到 DOM 元素。
	//btnEle 调用 addEventListener 为 btn 绑定点击事件 fibFunc。
	//在 fibFunc 中使用 numEle.Get("value") 获取到 numEle 的值（字符串），转为整型并调用 fib 计算出结果。
	//ansEle 调用 Set("innerHTML", ...) 渲染计算结果。
	done := make(chan int, 0)
	btnEle.Call("addEventListener", "click", js.FuncOf(fibFunc))
	<-done
}
