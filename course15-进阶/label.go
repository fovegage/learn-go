package main

func breakLabel() {
	//Loop:
	for i := 0; i < 100; i++ {
		if i == 5 {
			println("test", i)
			//  只能在for循环定义
			break
		}
	}

	println("start")
}

func gotoLabel1() {
	i := 1
	for i < 10 {
		if i == 5 {
			// 可以做异常跳转
			println(i)
			goto Loop
		}
		i++
	}
Loop:
	println("5")
}

func gotoLabel2() {
	i := 1
Loop:
	for i < 10 {
		if i == 5 {
			// 可以做异常跳转
			//i++
			// continue
			println(i)
			goto Loop
		}
		println(i)
		i++
	}
}

func main() {
	//gotoLabel1()
	gotoLabel2()
	//breakLabel()
}
