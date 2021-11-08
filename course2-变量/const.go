package main

const num1 int = 18

// 常量const必须进行初始化，因为它不能进行二次修改
func main() {
	const num2 int = 25
	println(num1, num2)
}
