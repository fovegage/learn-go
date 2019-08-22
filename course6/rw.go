package main

import "fmt"

func main() {
	var info = map[string]string{
		"name": "Gage",
		"sex":  "sex",
	}
	fmt.Println(info)

	// 字典元素读取
	fmt.Println(info["name"])

	// 字典元素修改
	info["name"] = "fovegage"
	fmt.Println(info)

	// 字典元素删除
	delete(info, "sex")
	fmt.Println(info)

	// 字典元素添加
	info["sex"] = "man"
	fmt.Println(info)

	// 访问不存在的key 返回空串
	fmt.Println(info["test"])

	// 如何判断key是否存在
	key_is_exist()
}

func key_is_exist() {
	var info = map[string]string{
		"name": "Gage",
		"sex":  "sex",
	}

	var vaule, ok = info["age"]
	if ok {
		fmt.Println(vaule)
	} else {
		fmt.Println("key is not exist")
	}
	info["age"] = "18"
	vaule, ok = info["age"]
	if ok {
		fmt.Println(vaule)
	} else {
		fmt.Println("key is not exist")
	}
}
