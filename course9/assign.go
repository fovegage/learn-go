package main

import "fmt"
import "unsafe"

type Point struct {
	x int
	y int
}

func main() {
	// 定义空接口，接收结构体
	var i interface{}
	var p = Point{
		x: 10,
		y: 10,
	}
	// 转换

	// 接口变量的赋值
	i = p
	fmt.Println(i, unsafe.Sizeof(i))

	var rx1 = i.(Point)
	//rx1.x = 20
	//rx1.y = 20
	p.x = 20
	p.y = 20
	fmt.Println(rx1, unsafe.Sizeof(rx1))

	// 指向指针的接口变量
	i = &p
	var rx2 = i.(*Point)
	//rx2.x = 30
	//rx2.y = 30
	p.x = 20
	p.y = 20
	fmt.Println(rx2, unsafe.Sizeof(rx2))
}
