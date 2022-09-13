package main

import (
	"fmt"
	"sync"
)

type User struct {
	name string
	age  int64
}


func main() {

	user := User{
		name: "Gage",
		age:  27,
	}
	//
	//fmt.Println(user.name)

	//fmt.Println(user.(*User).name)

	var res sync.Map

	res.Store("Gage", user)

	ss, ok := res.Load("Gage1")
	// https://studygolang.com/articles/21591?fr=sidebar
	// 这个是go语言转型
	//fmt.Println(ss, ok)
	fmt.Println(ss.(User).name, ok)

}
