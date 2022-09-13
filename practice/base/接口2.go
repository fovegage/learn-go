package main

import "fmt"

//定义一个鸭子接口
//Go 接口是一组方法的集合，可以理解为抽象的类型。它提供了一种非侵入式的接口。
//任何类型，只要实现了该接口中方法集，那么就属于这个类型。
type Duck interface {
	Gaga()
}

//假设现在有一个可达鸭类型
type PsyDuck struct{}

//假设现在有一个唐老鸭类型
type DonaldDuck struct{}

//可达鸭声明方法-满足鸭子会嘎嘎叫的特性
func (pd PsyDuck) Gaga() {
	fmt.Println("this is PsyDuck")
}

//唐老鸭声明方法-满足鸭子会嘎嘎叫的特性
func (dd DonaldDuck) Gaga() {
	//
	fmt.Println("this is DoningdDuck")
}

//要调用的函数 - 负责执行鸭子能做的事情,注意这里的参数,有类型限制为Duck接口
func DuckSay(d Duck) {
	d.Gaga()
}

// https://my.oschina.net/chinaliuhan/blog/3123026
// 核心 谁调用谁实现
func main() {
	//提示开始打印
	fmt.Println("duck typing")

	//实例化对象
	var pd PsyDuck    //可达鸭类型
	var dd DonaldDuck //唐老鸭类型

	//调用方法
	DuckSay(pd) //因为可达鸭实现了所有鸭子的函数,所以可以这么用
	DuckSay(dd) //因为唐老鸭实现了所有鸭子的函数,所以可以这么用
}
