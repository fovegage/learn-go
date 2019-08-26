package main

import "fmt"

// 定义接口
type Fruitable interface {
	eat()
}

// 定义结构体
// 匿名可以继承调用
type Fruit struct {
	name string
	Fruitable
}

// 关联

func (f Fruit) want() {
	fmt.Printf("I'm want eat")
	f.eat()
}

type Apple struct {
}

func (a Apple) eat() {
	fmt.Println("apple")
}

type Banana struct {
}

func (b Banana) eat() {
	fmt.Println("banana")
}
func main() {
	// 定义多态
	// 实现一个结构体调用多个方法
	var a = Fruit{"apple", Apple{}}
	var b  = Fruit{name:"banana", Fruitable:Banana{}}
	a.want()
	b.want()
}
