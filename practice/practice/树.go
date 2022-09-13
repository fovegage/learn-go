package main

import "fmt"

// 结构体 就是自定义的类型
// 不支持继承和多态
// 仅支持封装
type treeNode struct {
	Left, Right *treeNode
	Value       int
}

func (node *treeNode) print() {
	fmt.Println(node.Value)
}

// 结构体方法
func (node *treeNode) loop() {
	if node == nil {
		return
	}
	node.Left.loop()
	node.print()
	node.Right.loop()

}
func main() {
	var node treeNode
	node = treeNode{
		Value: 3,
	}
	node.Left = &treeNode{}
	node.Right = &treeNode{
		Left:  nil,
		Right: nil,
		Value: 8,
	}

	fmt.Println(node)

	node.loop()
}
