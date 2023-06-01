package main

import "fmt"

func main() {
	//
	// 创建指针字符串
	str := new(string)
	*str = "hello"

	// 修改字符串内容
	*str = *str + " world"

	// 输出字符串
	fmt.Println(*str)
}
