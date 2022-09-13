package main

import (
	"fmt"
	"math"
)

func tr() {
	// go语言必须进行强制类型转换 无隐式转换
	// 注意 sqrt需要float 这里需要进行强制的转换  才可以
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}
func main() {
	tr()
}
