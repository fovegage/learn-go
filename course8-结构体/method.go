package main

import "fmt"
import "math"

type Circle struct {
	x      int
	y      int
	Radius int
}

// 面积
func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func main() {
	var c Circle = Circle{Radius: 5}
	fmt.Println(c.Area())

	// 如果修改一些值 用指针 普通赋值无法进行修改
	var p = &c
	fmt.Println(p.Area())
}
