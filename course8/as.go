package main

import "fmt"
import "unsafe"

type ArrayStruct struct {
	value [4]int
}

type SliceStruct struct {
	value []int
}

func main() {
	var as = ArrayStruct{[...]int{1, 2, 3, 4}}
	var ss = SliceStruct{value: []int{1, 2, 3, 4, 5}}
	// sizeof 4*8
	fmt.Println(as, unsafe.Sizeof(as))
	// sizeof {pointer, len, cap} 始终都是 3*8
	fmt.Println(ss, unsafe.Sizeof(ss))
}
