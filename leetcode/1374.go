package leetcode

import "strings"

func main() string {

	n := 11
	if n%2 == 1 {
		return strings.Repeat("a", n)
	} else {
		return strings.Repeat("a", n-1) + "b"
	}

}
