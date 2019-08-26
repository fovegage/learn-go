package main

import (
	"fmt"
	"unsafe"
)
// 接口变量只包含两个指针字段, 一个用来存储数据内存，一个用来存储接口的类型信息和对象的类型信息
func main() {
	var i interface{}
	var array  = [4]int{1, 2, 3 ,4}
	i = array

	// 8 * 2
	fmt.Println(unsafe.Sizeof(i))

	// 8 * 4
	fmt.Println(unsafe.Sizeof(array))

	// 8 * 2
	// 显示的继承了接口
	fmt.Println(unsafe.Sizeof(i))
}
