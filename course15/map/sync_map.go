package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// 存储键值对
	m.Store("signhub.algorithm.jingdong.jd_sign", "John")
	m.Store("age", 30)

	// 根据键获取值
	if value, ok := m.Load("signhub.algorithm.jingdong.jd_sign"); ok {
		fmt.Println("Name:", value)
	}

	// 删除键值对
	m.Delete("age")

	// 遍历映射中的所有键值对
	//m.Range(func(key, value interface{}) bool {
	//	fmt.Printf("Key: %v, Value: %v\n", key, value)
	//	return true // 返回 true 继续遍历，返回 false 则停止遍历
	//})
}
