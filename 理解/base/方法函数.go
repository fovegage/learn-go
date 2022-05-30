package main

import "fmt"

// 不传参数
func print() {
	fmt.Println("hello world")
}

// 传参数
func Add(a, b int) {
	fmt.Println(a + b)
}

// 传参数  有返回值
func Sub(a int, b int) int {
	return a - b
}
func main() {
	print()
	Add(3, 4)
	fmt.Println(Sub(5, 9))
}
