package main

/*
函数式编程：https://learnku.com/articles/59902

函数作为go语言的一等公民
*/

func adder5050() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	// 匿名函数
	a := adder5050()
	for i := 0; i <= 100; i++ {
		// 每次循环匿名函数值会被临时保存
		// 即闭包
		println(a(i))
	}
}
