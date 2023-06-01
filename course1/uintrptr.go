package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 结合c++指针
	p := &Person{Name: "Gage", Age: 21}
	fmt.Println(uintptr(unsafe.Pointer(p)))
	//fmt.Println(uintptr(p))
}
