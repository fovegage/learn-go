package main

import "fmt"

func main() {
	var s1 []int = make([]int, 5, 8)
	for i := 0; i < len(s1); i++ {
		s1[i] = i + 1
	}
	fmt.Println(s1, len(s1), cap(s1))

	var c []int = make([]int, 2, 5)
	// 把 s1 的前两位拷贝到 c
	var n = copy(c, s1)
	fmt.Println(n, c)
}
