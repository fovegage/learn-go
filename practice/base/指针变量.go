package main

import "fmt"

// java和python 一般都是引用传递
// python 的引用传递区分在 可变类型和不可变类型上面
// go语言只有值传递  使用指针相当于值传递
// 值传递 copy了值   引用传递 引用的内存地址

func byVal(a, b int) (int, int) {
	// 值传递不会改变
	// 引用传递 就是说 a赋给了b  b赋给了a  这里是值传递
	b, a = a, b // 如果是引用传递的话  a赋给了b  b赋给了a 这里不是 // a=3 b=4 //作用域函数内部，不影响外部
	fmt.Println(a, b)
	return b, a
}
func byVal2(a, b int) (int, int) {
	// 值传递不会改变
	// 引用传递 就是说 a赋给了b  b赋给了a  这里是值传递
	b, a = a, b // 如果是引用传递的话  a赋给了b  b赋给了a 这里不是 // a=3 b=4 //作用域函数内部，不影响外部
	return a, b
}
func byVal1(a, b int) (int, int) {
	return b, a
}
func byRef(a , b *int) {
	*b, *a = *a, *b // 作用域全局
}
func main() {
	a, b := 3, 4
	////fmt.Println(byVal2(a, b))
	v1, v2 := byVal(a, b)  // 这个只可以在内部进行交换
	fmt.Println(v1, v2)
	fmt.Println(a, b)
	// &a 取a的地址传进去

	// 指针是一个全局的状态  修改后全局也就修改了
	byRef(&a, &b) // * 代表指针内存地址  & 代表取出内存地址的数据
	fmt.Println(a, b)

	//
}
