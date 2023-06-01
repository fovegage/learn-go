package main

import "fmt"

type Circle struct {
	x int
	y int
}

func main() {
	var c1 Circle = Circle{x: 20, y: 30}
	var c2 = c1
	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)
	// 注意 结构体内容 不共享 浅拷贝
	c2.x = 100
	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)

	var c3 *Circle = &Circle{x: 10, y: 15}
	var c4 *Circle = c3
	fmt.Printf("%+v\n", c3)
	fmt.Printf("%+v\n", c4)

	// 指针创建 结构体内容共享 只拷贝指针地址
	c3.x = 1000
	fmt.Printf("%+v\n", c3)
	fmt.Printf("%+v\n", c4)
}
