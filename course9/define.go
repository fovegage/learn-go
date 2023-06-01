package main

// 利用接口指向结构体变量
import "fmt"

// 定义接口
type Smellable interface {
	smell()
}

type Eatable interface {
	eat()
}

// 定义结构体
type Apple struct{}

// 关联结构体和接口
func (a Apple) smell() {
	fmt.Println("apple can smell")
}

func (a Apple) eat() {
	fmt.Println("apple can eat")
}

// 定义结构体
type Flower struct{}

// 关联接口和结构体
func (f Flower) smell() {
	fmt.Println("flower can smell")
}

func main() {
	var s1 Smellable
	var e1 Eatable
	var apple = Apple{}
	var flower = Flower{}

	s1 = apple
	s1.smell()

	s1 = flower
	s1.smell()

	e1 = apple
	e1.eat()
}
