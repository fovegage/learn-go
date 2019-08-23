package main

import "fmt"

func main() {
	var s1 = "hello" // 静态字面量
	var s2 = ""
	for i := 0; i < 10; i++ {
		s2 += s1 // 动态构造
	}
	fmt.Println(s1, len(s1))
	fmt.Println(s2, len(s2))

	// 注意只可访问，不可修改，和Python一致
	fmt.Printf("%c", s1[3])
}
