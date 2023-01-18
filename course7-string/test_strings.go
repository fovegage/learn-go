package main

import (
	"fmt"
	"strings"
)

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
func main() {
	fmt.Print(FirstUpper("hello"))
}
