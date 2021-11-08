package main

import (
	"errors"
	"fmt"
)

func main() {
	var err1 = errors.New("some error")
	fmt.Println(err1)

	//定制输出
	var e = "some error"
	var err2 = fmt.Errorf("Gage: %s", e)
	fmt.Println(err2)
}
