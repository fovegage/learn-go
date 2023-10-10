package label

import "fmt"

func main() {
	// 使用标签和break语句来控制循环
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break outerLoop // 跳出outerLoop标签所指定的循环
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}

	fmt.Println("\n使用goto语句跳转到标签处的代码块：")

	// 使用goto语句跳转到标签处的代码块
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				goto printResult // 跳转到printResult标签处的代码块
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}

printResult:
	fmt.Println("跳转到了标签处的代码块")
}
