package main

// 注意 ’‘ “” 这里和c语言一直
func main() {
	// 字符串类型
	var name string = "Gage"

	// int 32位机器4个字节 64位机器8个字节
	// int 表示正负数 uint 表示非负数
	var num1 int = 25
	var num2 uint = 25

	// 浮点数
	var h float32 = 3.14
	var f float64 = 3.1415926

	// 布尔类型
	var sex bool = true

	// 字节类型
	var b byte = 'a'

	// 指针类型
	var value int = 5
	var pointer *int = &value

	println(name, num1, num2, b, sex, h, f, pointer, *pointer)
}
