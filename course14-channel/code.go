package main

import (
	"fmt"
	"time"
)

func modeOne() {
	timer := time.NewTimer(time.Second * 5)
	<-timer.C
	fmt.Println("mode one: Time out!")
}

func modeTwo() {
	select {
	case <-time.After(time.Second * 5):
		fmt.Println("mode two: Time out!")
	}

}

/*
// 创建一个整型通道
ch := make(chan int)
// 尝试将0通过通道发送
ch <- 0

-----
# 接收数据
data := <-ch


*/

func main() {
	// 如何理解
	// http://c.biancheng.net/view/97.html
	//<-app.exitChan
	//app.cancel()
	//depInj.Stop(app.ctx)

	modeOne()
	modeTwo()
}
