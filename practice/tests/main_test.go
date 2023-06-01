package main

// 需要符合  xx_test 才可以被识别到

// 对方法测试的时候 则必须是Test__开头

// go test 只打印基本的信息
// go test -v 打印更加详细的信息
import (
	"fmt"
	"testing"
)

func testPrint1(t *testing.T) {
	t.SkipNow() // 可以忽略错误 直接运行
	fmt.Printf("hello t1")
}

// subtest 测试

func testPrint2(t *testing.T) {
	fmt.Printf("this is t2")
}

// 由于 go test 不能保证顺序 因此可以用下面的形式控制  子测试的话 将不在使用 Test 大写开头  而是 test_xx 小写开头
func TestAll(t *testing.T) {
	t.Run("t1", testPrint1)
	t.Run("t2", testPrint2)

}

// 特殊的测试文件  名字就这样写
func TestMain(m *testing.M) {
	fmt.Printf("test being")
	m.Run() // 必须调用   这儿调用的是 TestAll

}

// beachmark测试

func hello(b int) int {
	//fmt.Println(b)
	return b
}

// Beachmark  开头写法固定
// 使用 go test -bench = . 进行触发
func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello(i)
	}

}
