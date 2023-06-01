package main

import (
	"fmt"
	"time"
)

// 注意代码形式上虽然是协程套协程，但是他们是并行执行的
// 只有一个主协程，其他子协程都是平行的
// 主协程不会等到子协程运行结束后再停止
// 主协程一旦结束，其他协程立即消亡
func main() {
	fmt.Println("main goroutine start")
	go func() {
		fmt.Println("child goroutine start")
		go func() {
			fmt.Println("child child goroutine start")
			go func() {
				defer func() {
					// 为了防止子协程的异常向主协程传播，应该用recover()进行处理
					if err := recover(); err != nil {
						panic("人为异常")
					}
					fmt.Println("child child child goroutine start")

				}()
			}()
		}()
	}()
	time.Sleep(time.Second)
	fmt.Println("main goroutine end")
}
