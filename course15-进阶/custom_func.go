package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
自定义函数
fib: 逐项相加
1, 1, 2, 3, 5, 8
*/

// IntGen 定义新的数据类型
type IntGen func() int

func (ig IntGen) Read(p []byte) (n int, err error) {
	next := ig()
	if next > 100 {
		return 0, io.EOF
	}
	// Sprintf 理解
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func fib() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		// 0 1
		// 1 1
		// 1 2
		// 2 3
		// 3 5
		// 5 8
		return a
	}
}

func printFib(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		println(scanner.Text())
	}

}

func t1(a any) {

}

func main() {
	f := fib()
	printFib(f)
}
