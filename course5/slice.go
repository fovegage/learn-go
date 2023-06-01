package main

import "fmt"

//  切片一般包含 [start, end], 若不指定start,则从起始位置开始，若不指定end,
func main() {
	var s1 []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s1, len(s1), cap(s1))

	// [1 2 3 4 5] 5 7  未指定start
	var s2 = s1[:5]
	fmt.Println(s2, len(s2), cap(s2))

	//[4 5] 2 4 指定了start和end   cap = len-start
	var s3 = s1[3:5]
	fmt.Println(s3, len(s3), cap(s3))

	// [2 3 4 5 6 7] 6 6 未指定start
	var s4 = s1[1:]
	fmt.Println(s4, len(s4), cap(s4))

	// [1 2 3 4 5 6 7] 7 7 相当于浅拷贝
	var s5 = s1[:]
	fmt.Println(s5, len(s5), cap(s5))

	// s5 和 s6 是等效的，均是浅拷贝
	var s6 = s1
	fmt.Println(s6, len(s6), cap(s6))
}
