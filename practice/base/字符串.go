package main

import (
	"fmt"
	"unicode/utf8"
)

func loop(s2 []byte) {
	for len(s2) > 0 {
		ch, size := utf8.DecodeRune(s2) // 取第一个
		s2 = s2[size:]
		fmt.Println(s2, ch, size)
		fmt.Printf("%c ", ch)
	}
}

// %X 16进制
// %c 字节
// %s 字符串
// %d 数字
// 英文1 中文3
// 具体参见：https://blog.csdn.net/qq_22660775/article/details/89487263
func main() {
	s1 := "hello你好world" // utf-8
	// 直接运行 aci码
	for _, v := range s1 {
		//fmt.Println(k, v)
		//fmt.Print(v)
		fmt.Printf("%d:", v)
		fmt.Printf("%c ", v)
	}
	fmt.Println()
	// 16进制  默认unicode输出
	for _, v := range s1 {
		//fmt.Println(k, v)
		fmt.Printf("%X ", v)
	}
	fmt.Println()
	// byte 为16 把 中文拆分了三个  utf-8
	for _, v := range []byte(s1) {
		//fmt.Println(v)
		fmt.Printf("%X ", v)
	}
	fmt.Println()
	// rune  unicode编码
	for k, v := range []rune(s1) {
		fmt.Printf("%d:%c ", k, v)
	}
	fmt.Println()
	// 字符数量
	fmt.Println(len(s1), utf8.RuneCountInString(s1), len([]rune(s1)))
	// 底层输出
	s2 := []byte(s1)

	fmt.Printf("%c", 20320)
	fmt.Println()
	//loop(s2)

	fmt.Print(len(s1))
	fmt.Println()

	fmt.Println(s2[1:])
	//fmt.Println(utf8.DecodeRune(s2))

	fmt.Print("\xf0\x9f\x98\x81")
}
