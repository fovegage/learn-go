package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" name:"张三"`
	Age  int
}

func (_ User) Say(name string, num int) (string, string) {
	fmt.Println("user 说话", name, num)
	return "hello", "world"
}

func main() {
	// https://learnku.com/articles/65666
	println("start")
	user := User{Name: "张三", Age: 10}
	res := reflect.ValueOf(user).MethodByName("Say").Call([]reflect.Value{reflect.ValueOf("该说话了"), reflect.ValueOf(1)})
	println(res[0].Interface().(string))

	println("----------")

	api := &User{}
	methodFun := reflect.TypeOf(api)
	for i := 0; i < methodFun.NumMethod(); i++ {
		method := methodFun.Method(i)
		println(method.Name)
	}
}
