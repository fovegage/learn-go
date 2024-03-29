package main

import "fmt"

/*
闭包
*/
func c15Test1() {
	a := func() {}

	println(a)
}

func adder() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	c15Test1()

	// 函数作为返回值
	a := adder()
	fmt.Println(a(1)) //第一次调用时sum = 0 ， sum = 0 + 1，返回值为1
	fmt.Println(a(2)) //第二次调用时sum = 1 ，sum = 1 +2，返回值为3
	fmt.Println(a(9)) //第三次调用时sum = 9，sum = 3 + 9，返回值为12
}
