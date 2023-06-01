package main

import "fmt"

func main() {
	// 中文
	//word := "截取字符串"
	word := "hello"
	println(word[:3])

	// byte
	println([]byte("hello"))
	println([]byte("你好"))

	// https://www.runoob.com/w3cnote/byte-character.html
	// todo: 结合之前的python字符
	// rune byte
	// byte 一个字节(8位) uint8   2^8
	// rune 四个字节(32位) int32  2^32
	// UTF-8 编码中，一个英文字为一个字节，一个中文为三个字节。
	first := "社区"
	fmt.Println([]rune(first))
	fmt.Println([]byte(first))

	// english
	two := "hello"
	fmt.Println([]rune(two))
	fmt.Println([]byte(two))
}
