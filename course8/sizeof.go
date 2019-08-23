package main

import "fmt"
import "unsafe"

type Circle struct {
	x int
	y int
	z int
}

func main() {
	var c1 Circle = Circle{x: 20}
	fmt.Printf("%+v\n", c1)
	// 64位机器  一个int为8  3*8=24
	fmt.Println(unsafe.Sizeof(c1))
}
