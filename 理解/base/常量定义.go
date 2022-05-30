package main

import (
	"fmt"
	"math"
)

// 定义常量 不会进行爆红  而变量会
func simple() {
	const name = "abc.txt"
	const age1, age2 int = 23, 25
	const (
		t1, t2 = 23, true
	)
	fmt.Println(name, age1, age2, t1, t2)
}

func tr1() {
	const a, b = 3, 4
	var c float64
	c = math.Sqrt(a*a + b*b) // 自动判断 无需转型
	fmt.Println(c)
}

func enums() {
	// 定义枚举类型 使用 ioto后会自动进行自增
	const (
		python = 0
		java   = 1
		c      = 2
		golang = 3
	)
	fmt.Println(python)

	const (
		javascript = iota
		ruby
	)
	fmt.Println(ruby)
}
func main() {
	simple()
	tr1()
	enums()
}
