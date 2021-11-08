package main

import "fmt"

// 定义结构体

type Circle struct {
	x      int
	y      int
	Radius int
}

func main() {
	// 初始化结构体
	// 可以部分指定，为赋值的会自动显示类型初始值
	var c1 Circle = Circle{
		x:      100,
		y:      200,
		Radius: 7,
	}
	fmt.Printf("%+v\n", c1)

	// 默认类型初始值
	var c2 Circle = Circle{}
	fmt.Printf("%+v\n", c2)

	// 如果是下面的方式创建，则结构体的值必须全部提供
	var c3 Circle = Circle{10, 20, 3}
	fmt.Printf("%+v\n", c3)

	// 指针形式, & 表示结构体是指针类型
	var c4 *Circle = &Circle{10, 20, 5}
	fmt.Printf("%+v\n", c4)

	// new()进行创建,构建零值结构体，返回类型为指针
	var c5 *Circle = new(Circle)
	fmt.Printf("%+v\n", c5)

	// 定义初始化
	var c6 Circle
	fmt.Printf("%+v\n", c6)

	// nil初始化,这样创建的结构体只占一个字节，并未指向真实的内存
	var c7 *Circle = nil
	fmt.Printf("%+v\n", c7)
}
