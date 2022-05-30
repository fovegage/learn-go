package main

import "fmt"

// []int 代表切片
func testA(arr []int) {
	arr[0] = 100
}

// 数组长度 数组容量
func testB(arr []int) {
	// cap计算 原始数组长度 - start(数组开始游标)
	s1 := arr[2:6]
	fmt.Println(s1, len(s1), cap(s1))
	// s1的len 3 4 5 6
	// s1的cap 3 4 5 6 7 8 9 10

	// s2 在 s1 cap的基础上进行slice
	// s2 的len 7 8 9 s2的cap 7 8 9 10
	// 注意这儿 如果超过数组的长度 会在cap()长度进行截取 如果超过cap()则会报错
	s2 := s1[4:7]
	fmt.Println(s2, len(s2), cap(s2))

	// 切片完 变量就重新开辟内存空间了
	s3 := s1[:4]
	fmt.Println(s3, len(s3), cap(s3))

	// append 元素
	// 当append元素的时候 golang会自动的扩充cap
	s4 := append(s3, 10)
	fmt.Println(s4)
	s5 := append(s4, 11)
	fmt.Println(s5)
	s6 := append(s5, 12)
	fmt.Println(s6)
	s7 := append(s6, 13)
	fmt.Println(s7, len(s7), cap(s7))
	s8 := append(s7, 14)
	fmt.Println(s8, len(s8), cap(s8))
}

func createSlice() {
	a1 := make([]int, 16)
	fmt.Println(a1, cap(a1))

	a2 := make([]int, 16, 32)
	fmt.Println(a2, cap(a2))
}

func copyArr() {
	a1 := []int{2, 3, 4, 5}
	fmt.Println(a1)
	a2 := make([]int, 10)
	//copy(a1, a2)
	//fmt.Println(a1, "tests")
	copy(a2, a1) // 把a1拷贝到a2
	fmt.Println(a2)
	// 删除元素 这样就删除了a2[3]
	a3 := append(a2[:3], a2[4:]...)
	fmt.Println(a3)

}
func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	testA(arr[2:])
	fmt.Println(arr)
	// 这里的切片和python的一致
	fmt.Println(arr[:2])
	// 全量传入
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testB(arr1[:])

	// make 创建slice
	createSlice()
	// 数组拷贝
	copyArr()
}
