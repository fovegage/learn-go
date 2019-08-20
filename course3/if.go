package main

func main() {
	println(max(5, 5))
}

// if 后面不需要加 () go语言没有三元表达式，只能通过 if...else 实现
func max(a int, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	}
	return 0
}
