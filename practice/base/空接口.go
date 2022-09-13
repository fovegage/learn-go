package main

import "fmt"

//  空接口可以接收任意的数据类型
func printAll(vals []interface{}) {

	for key, val := range vals {
		fmt.Println(key, val)
	}

}

func main() {
	names := []string{"Gage", "Tony"}
	vals := make([]interface{}, len(names))
	for key, val := range names {
		vals[key] = val
	}
	printAll(vals)

}
