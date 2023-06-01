package main

import "fmt"

func main() {
	// make(type, length, capacity)

	// 零值切片
	// 使用make函数创建切片，长度为5，容量为8
	var s1 []int = make([]int, 5, 8)
	fmt.Println(s1)
	// 满容量数组
	var s2 []int = make([]int, 8)
	fmt.Println(s2)
	// 推导创建
	s3 := make([]int, 5, 8)
	fmt.Println(s3)

	// 满容切片
	// 切片初始化
	s4 := []int{1, 2, 3, 4, 5}
	fmt.Println(s4, len(s4), cap(s4))

	// nil 切片
	var s5 []int
	fmt.Println(s5, len(s5), cap(s5))

	//空切片
	var s6 []int = []int{}
	fmt.Println(s6, len(s6), cap(s6))

	var s7 []int = make([]int, 0)
	fmt.Println(s7, len(s7), cap(s7))
}
