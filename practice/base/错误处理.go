package main

import "fmt"

func errHandle() {
	// 加了defer可以按栈顺序输出  自动回收异常
	defer fmt.Println(110)
	defer fmt.Println(120)
	fmt.Println(2)
	panic("tests error")
	fmt.Println(3)
}
func main() {
	errHandle()
}
