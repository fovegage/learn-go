package main

import (
	"fmt"
)

// 在函数外部，不可以直接使用 := 进行命名
// 连续定义多个变量 var()
// 无全局变量，包内引用
// 可以不赋初值，会默认使用初始值 eg: int=0 string=''
var name string = "Gage"

var (
	t1 = 25
	t2 = "Gage"
)

func varZero() {
	var first string = "Hello world" // 定义数据类型
	var name = "hello"               // 可以不定义类型，自动判断
	var one, two = "hi", true        // 连续定义
	//var t3 string, t4 int // 这样定义是不对的，自动判断即可
	three, four := 26, "hi"
	age := 25 // 可以不使用var
	fmt.Println(name, age, first, one, two, three, four)
}
func printNullValue() {
	// printf 格式化输出  println 打印值
	var name string
	var age int
	fmt.Printf("%q, %d\n", name, age)
}
func main() {
	varZero()
	printNullValue()
	fmt.Println(name)
	fmt.Println(t1, t2)
}
