package main

import "fmt"

// 定义结构体  go语言没有继承和多态
// 这里可以理解为go语言的面向对象  为其添加类变量
// 就是 自己新整合的数据类型
type Books struct {
	title string
	sn    int32
}

func (book Books) print() {
	fmt.Println(book)
}

//
func (book *Books) printByRef() {
	book.title = "Java"
	fmt.Println(book.title)
}
func main() {
	python := Books{
		title: "Python CookBook",
		sn:    81919,
	}
	// 输出
	fmt.Println(python)
	// 访问属性
	fmt.Println(python.title)
	// 修改属性值
	python.title = "Golang"
	fmt.Println(python)
	// 结构体方法  值传递
	python.print()
	// 引用传递
	python.printByRef()
}
