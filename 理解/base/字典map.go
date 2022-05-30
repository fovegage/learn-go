package main

import "fmt"

func main() {
	m1 := map[string]string{
		"name": "Gage",
		"sex":  "man",
		"high": "185",
	}

	m2 := make(map[string]int) // 值默认为空

	m3 := map[string]int{} // 值默认为 nil

	fmt.Println(m1, m2, m3)

	// 遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 取值
	v1 := m1["name"]
	fmt.Println(v1)

	// 判断key是否存在 不存在取为空
	v2, ok := m1["sex"]
	// ok 接收 是否存在 true false
	fmt.Println(v2, ok)

	// 删除元素
	delete(m1, "sex")
	// 不可以使用 :=  上面已经定义过  这里赋值修改即可
	v2, ok = m1["sex"]
	fmt.Println(v2, ok)
}
