package main

import "fmt"

// 应该注意的是：数组和切片的区别
func main() {
	var s []int = []int{1, 2, 3, 4, 5}
	for index := range s {
		fmt.Println(index, s[index])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
