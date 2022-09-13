package main

import (
	"fmt"
	"io/ioutil"
)

func fetch() {
	const filename = "E:/learn/golang/testfiles/fetch.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}
}

func fetch1() {
	// 等于 nil  说明没有报错
	// 先运行; 前面的   然后在走逻辑
	// 就是赋值了
	const filename = "E:/learn/golang/testfiles/fetch.txt"
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Printf("%s\n", contents)
	} else {
		fmt.Println(err)
	}
}

func fetch2(a, b int, c string) int {
	var result int
	switch c {
	case "+":
		result = a + b
	case "-":
		result = a - b
	default: // 其他值判断
		fmt.Printf("%s", "error string")
	}
	return result
}

func fetch4(a int) string {
	switch {
	case a < 0 || a > 100:
		// python raise   将中断程序
		panic(fmt.Sprintf("%s", "self define except"))
	case a < 29:
		return "F"
	default:
		return "A"
	}
}
func main() {
	fetch()
	fetch1()
	fmt.Println(fetch2(1, 2, "+-"))
	//fmt.Println(fetch4(691))
}
