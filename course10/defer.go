package main

import (
	"fmt"
	"os"
)

// 当多个defer出现的时候，defer执行顺序是和书写顺序相反的
func main() {
	f1, err := os.Open("catch.go")
	if err != nil{
		fmt.Println("file1 is not exist")
	}
	defer func() {
		fmt.Println("close1 catch")
		f1.Close()
	}()

	f2, err := os.Open("error.go")
	if err != nil{
		fmt.Println("file2 is not exist")
	}
	defer func() {
		fmt.Println("close2 error")
		f2.Close()
	}()

}
