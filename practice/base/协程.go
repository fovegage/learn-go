package main

import "fmt"

func main() {
	//arr := [1000]int{}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			//arr[i]++
			fmt.Printf("print %v \n", i)
		}(i)
	}
	//fmt.Print(arr)
}
