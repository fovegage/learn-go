package main

import (
	"bufio"
	"fmt"
	"os"
)

// go语言的 分支和循环语句均不带括号
// 只要方法带括号
func sum() {
	sm := 0
	for i := 1; i <= 100; i++ {
		sm += i
	}
	fmt.Println(sm)
}

// go语言没有while  但是for 可以省略起始条件和递增条件  只剩条件表达式 就是go语言的while

func printFile() {
	//
	contents, err := os.Open("E:/learn/golang/testfiles/fetch.txt")
	if err != nil {
		panic("error")
	} else {
		scanner := bufio.NewScanner(contents)
		// 相当于 while 只有条件表达式
		// scanner.Scan() 判断是否还有内容
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

}

func forever() {
	// 死循环
	for {
		fmt.Println("dead while")
	}
}
func main() {
	sum()
	printFile()
	forever()
}
