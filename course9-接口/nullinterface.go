package main

import "fmt"

func main() {
	// interface是一个万能容器，可以接收任何数据类型
	var info = map[string]interface{}{
		"name": "Gage",
		"age":  25,
		"sex":  true,
	}
	fmt.Println(info)
	// 类型转换
	var name = info["name"].(string)
	fmt.Println(name)
	var age = info["age"].(int)
	fmt.Println(age)
	var sex = info["sex"].(bool)
	fmt.Println(sex)
}
