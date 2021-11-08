package main

import "fmt"

type Circle struct {
	x int
	y int
}

func ByValue(c Circle) {
	c.x *= 2
}

func ByPointer(c *Circle) {
	c.x *= 2
}
func main() {
	var c1 Circle = Circle{x: 50, y: 100}
	//无法修改
	ByValue(c1)
	fmt.Printf("%+v\n", c1)

	// 可以修改
	ByPointer(&c1)
	fmt.Printf("%+v\n", c1)
}
