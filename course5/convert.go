package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	var s = a[2:5]
	fmt.Println(s, len(s), cap(s))
	s[2] = 888
	fmt.Println(s, len(s), cap(s))

	// 注意切片的操作会影响到数组,修改数组也会影响到切片
	// 他们两者之间是相互影响的
	fmt.Println(a, len(a))
}
