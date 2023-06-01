package main

import (
	"fmt"
	"os"
)

func main() {
	var f, err = os.Open("E:\\go\\course10-异常\\error.go")
	if err != nil {
		println(err.Error())
		return
	}
	// 相当于Java或Python的finally
	defer f.Close()

	// 字节切片存储内容
	var content = []byte{}
	// 缓冲
	var buf = make([]byte, 100)
	// 相当于死循环
	for {
		n, err := f.Read(buf)
		if n > 0 {
			content = append(content, buf[:n]...)
		}
		if err != nil {
			break
		}
		fmt.Println(string(content))
	}
}
