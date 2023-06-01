package main

// go语言没有静态变量
// 如果全局变量的首字母大写则外包可以访问，如果小写则只能在本包内调用
var global string = "Gage"

// 全局变量和局部变量
func main() {
	var local string = "Gage"
	println(global, local)
}
