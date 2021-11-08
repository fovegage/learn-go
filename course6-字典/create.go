package main

import "fmt"

func main() {
	// make创建
	var m map[int]string = make(map[int]string, 16)
	fmt.Println(m, len(m))

	// 初始化 注意 key的类型为int;value的类型为string
	var m1 map[int]string = map[int]string{
		90: "优秀",
		80: "良好",
	}
	fmt.Println(m1)

	// 自动推导
	var info = map[string]string{
		"name": "Gage",
		"sex":  "sex",
	}
	fmt.Println(info)
}
