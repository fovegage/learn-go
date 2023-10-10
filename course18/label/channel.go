package label

import "fmt"

func ChanTest() {
	ch := make(chan int)

	// 启动一个goroutine来模拟计算乘积并发送到通道
	go func() {
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				product := i * j
				if product > 4 {
					ch <- product // 将乘积发送到通道
				}
			}
		}
		close(ch) // 关闭通道，表示没有更多数据发送
	}()

	fmt.Println("使用通道接收数据：")
	// 从通道接收数据
	for product := range ch {
		fmt.Printf("Received: %d\n", product)
	}

	fmt.Println("完成")
}

func TestChan1() {
	reqs := make(chan string, 10)

	reqs <- "https://www.baidu.com"
	go func() {
		for {
			urls := []string{"11", "22"}
			for _, url := range urls {
				reqs <- url
			}
		}
		close(reqs) // 关闭通道，通知发送者不再发送数据
	}()

	for url := range reqs {
		println(url)
	}
}
