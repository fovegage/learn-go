package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		m.Store("key", "value from goroutine 1")
	}()

	go func() {
		defer wg.Done()
		m.Store("key", "value from goroutine 2")
	}()

	wg.Wait()

	// 获取并打印存储在 "key" 上的值
	if value, ok := m.Load("key"); ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found.")
	}
}
