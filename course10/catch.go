package main

import (
	"fmt"
)

var Err = fmt.Errorf("must >= 0")

func test(num int) int {
	// 栈捕捉
	if num < 0 {
		panic(Err)
	}
	return 1
}
func main() {
	// 可以不输出栈异常信息
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(test(2))
	fmt.Println(test(-5))
	// 异常发生后，不会执行后面的
	fmt.Println(test(1))
}
