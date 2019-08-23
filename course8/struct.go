package main

import "fmt"

type Point struct {
	x int
	y int
}

// 结构体嵌套 非匿名
type Circle1 struct {
	p      Point
	Radius int
}

// 结构体嵌套 匿名
type Circle2 struct {
	Point
	Radius int
}

// 将作为结构体的一个方法
func (p Point) show() {
	fmt.Println(p.x, p.y)
}
func name() {
	var c = Circle1{
		p:      Point{x: 10, y: 20},
		Radius: 5,
	}

	fmt.Printf("%+v\n", c.p)
	fmt.Printf("%+v\n", c.p.x)
	fmt.Printf("%+v\n", c)
}

func unname() {
	var c = Circle2{
		Point:  Point{x: 1, y: 2},
		Radius: 5,
	}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", c.x)
	fmt.Printf("%+v\n", c.y)
	fmt.Printf("%+v\n", c.Point)
	c.show()
	c.Point.show()
}
func main() {
	// 非匿名结构体
	name()
	// 匿名结构体 可以对嵌套的结构体的变量和方法进行继承 而非匿名不可以
	unname()
}
