package main

import "fmt"

func main() {
	var s1 []int = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s1)

	// 满容数组，append后不共享底部数组
	// 由于 s1 的长度为6,则append后的 s2 长度为12
	var s2 = append(s1, 7)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	// 非满容数组共享底部变量
	var s3 = append(s2, 8)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	// 只扩容，不使用
	var s4 []int = []int{1, 2, 3}
	_ = append(s4, 6)
	fmt.Println(s4, len(s4), cap(s4))
}
