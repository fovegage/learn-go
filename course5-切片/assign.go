package main

import "fmt"

func main() {
	var s []int = make([]int, 5, 8)
	for i := 1; i < len(s); i++ {
		s[i] = i
	}
	fmt.Println(s)

	// 浅拷贝 共享底部底层数组
	var s1 = s
	fmt.Println(s1)
	s[3] = 666
	fmt.Println(s)
	fmt.Println(s1)
}
