package main

import "fmt"

func complexTest() {
	// int: int int8 int16 int32 int64
	// float: float32 float64
	// string: string
	// bool: true false
	// byte:
	// rune:
	var score complex64 = 1 + 3i
	fmt.Println(score)
}
func main() {
	complexTest()
}
