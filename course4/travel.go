package main

import "fmt"

func main() {
	travel()
}

func travel() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	for index := range a {
		fmt.Println(index, a[index])
	}
	for index, value := range a {
		fmt.Println(index, value)
	}
}
