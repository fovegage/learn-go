package main

import "fmt"

func arrTest(arr [5]int) {
	for _, v := range arr {
		fmt.Println(v)
	}

}
func arrTest1(arr [5]int) [5]int {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
	return arr
}

func main() {
	//数组的几种形式
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5} // 不声明var 必须给初值{}
	arr3 := [...]int{2, 4, 5}     // 自己数有几个值
	arr4 := [4][5]int{}           // 默认{} 初始值
	fmt.Println(arr1, arr2, arr3, arr4)

	// 对数组进行遍历
	// 类似于python的enumerate
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	// _ 和python 类似
	for _, v := range arr3 {
		fmt.Println(v)
	}

	// arrTest要求传入5个长度的数组 当传入长度为3的数组会报错
	arrTest(arr1)
	//arrTest(arr3)

	// 值传递即深拷贝 不会改变值 只在函数内部生效
	fmt.Println(arr1, arrTest1(arr1))

}
