package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic("create fail")
	}

	defer file.Close() // 执行后关闭

	write := bufio.NewWriter(file)
	defer write.Flush()  // 出栈顺序 先刷新在关闭
	for i:=1;i<20;i++{
		fmt.Fprintln(write, i)
	}

}
func main() {
 writeFile("E:/learn/golang/testfiles/t1.txt")
}
