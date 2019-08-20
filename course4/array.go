package main

import (
	"fmt"
)

// 默认生成长度为9的定长数组，初始值均为0
func main() {
	var a1 [9]int
	fmt.Println(a1)
	var a2 = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a2)
	var a3 [9]int = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a3)
	a4 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a4)
}
