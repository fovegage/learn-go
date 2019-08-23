package main

import "fmt"

var s1 = "你好hello"

func main() {
	//byte()
	str()
}

// 按字节遍历
func byte() {
	for i := 0; i < len(s1); i++ {
		//fmt.Println(s1[i])
		fmt.Printf("%x ", s1[i])
	}
}

// 按字符遍历
func str() {
	for pos, code := range s1 {
		fmt.Printf("%d %d ", pos, code)
	}
}
