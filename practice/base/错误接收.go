package main

import (
	"errors"
	"fmt"
)

// error panic
// 对于可预见的 error
// 不可预见的 panic

func errorTest() {
	// 匿名函数  直接调用
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("Expect error: %v", err)
		} else {
			panic(err)
		}
	}()

	panic(errors.New("self define expect error"))
	//panic(111) // 不是error类型  不可预知
}
func main() {

	errorTest()

}
