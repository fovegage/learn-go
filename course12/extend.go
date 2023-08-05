package main

type Animal interface {
	Say()
}

type Man1 struct {
	Animal
}

func (m *Man1) Say() {
	println("say")
}

// NewMan 通过集成实现  Man1相当于实现类  New 相当于构造函数
func NewMan() *Man1 {
	return &Man1{}
}

// NewMan 通过签名实现
//func NewMan() Animal {
//	return &Man1{}
//}

func main() {
	man := NewMan()
	man.Say()
}
