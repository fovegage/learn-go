package main

import "fmt"

func main() {
	// 衍生类型 rune 占用4个字节
	var s1 = 's'
	// 一个汉字占用三个字节 一个英文字母占用一个字节
	var s2 = "你好hello"
	fmt.Println(s1)
	fmt.Println(s2, len(s2))
}
