package main

import "fmt"

func main() {
	assign()
}

// 注意只能在两个数组长度相同时，才能进行copy操作
func assign() {
	var a [9]int = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b [9]int
	b = a
	a[0] = 111
	fmt.Println(b)
	fmt.Println(a)
}
