package main

import "fmt"

var m = map[int]string{
	90: "优秀",
	80: "良好",
}

func main() {

	fmt.Println(m)

	// 下面情况类似于Python的解压
	// 如果在判断key是否存在，[key, is_exist(boolean)]
	for key, value := range m {
		fmt.Println(key, value)
	}

	for key := range m {
		fmt.Println(key)
	}

	// 输出字典的key
	keys()
	// 输出字典的items
	items()
}

func keys() {
	// 空切片被创建后 在原本的长度有初始化的0值
	// 因此 下面的0 不能去掉 去掉会动态扩容
	var name = make([]int, 0, len(m))
	for key := range m {
		name = append(name, key)
	}
	fmt.Println(name)
}

func items() {
	var name = make([]string, 0, len(m))
	for key := range m {
		name = append(name, m[key])
	}
	fmt.Println(name)
}
