package main

import (
	"../pkg/tree"
	"fmt"
)

// 只有大写开头的才可以 被引入
// 如果扩展已有的类型  则在main中定义一个小写开头的私有方法
// 用& 引用的原因在于 不用& 只在 方法内有值  每调用一次会开辟一个新的内存空间
func main() {
	var node tree.Node
	node = tree.Node{
		Value: 3,
	}

	node.Left = &tree.Node{}
	node.Right = &tree.Node{
		Left:  nil,
		Right: nil,
		Value: 8,
	}

	fmt.Println(node)

	node.Loop()
}
