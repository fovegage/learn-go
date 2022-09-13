package main

import "fmt"

// https://learnku.com/go/t/35168
// http://c.biancheng.net/view/21.html
// & 这个取的是 内存地址
// * 这个取的是数据
func ref(val *string) {
	//fmt.Println(val)
	var key = "hello"
	val = &key
	fmt.Println(val)
}

func test1() {
	a := 7
	fmt.Println(&a)
	//
	fmt.Println(*&a)
	//
	b := &a
	*b++
	// 值在内存中存储着   修改完值回写到内存中
	fmt.Println(a, b, &a, &b, *b)
}

// *int 是一个数据类型声名 声名必须传入指针
// 在函数内部修改了值  调用函数接着读取到内存中的新值
func modifyRef(a *int) {
	*a++
	//fmt.Println(a)
}

// 只能在内部作用
func modifyVal(a int) {
	a++
}
func test2() {
	var a int
	//fmt.Println(&a)
	modifyRef(&a)
	//modifyVal(a)
	fmt.Println(a)
}
func main() {
	//var name string
	//ref(&name)
	//fmt.Println(name)
	test2()

}
