package main

func main() {
	//loop()
	hello()
}

func loop() {
	for {
		println("死循环")
	}
	// 等效
	//for true {
	//	println("死循环")
	//}
}

func hello() {
	for i := 0; i < 10; i++ {
		println("hello world")
	}
}
