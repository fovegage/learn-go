package main

import "fmt"

func main() {
	visit()
}

func visit() {
	var a [9]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a)

}
