package tree

import "fmt"

// 注意命名必须为大写 不然导不出去
// 结构体 就是自定义的类型
// 不支持继承和多态
// 仅支持封装
type Node struct {
	Left, Right *Node
	Value       int
}

func (node *Node) print() {
	fmt.Println(node.Value)
}

// 结构体方法
func (node *Node) Loop() {
	if node == nil {
		return
	}
	node.Left.Loop()
	node.print()
	node.Right.Loop()

}
