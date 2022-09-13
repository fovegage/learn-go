package main

type Person struct {
}

func (p Person) eat() {
	print("parent eat\n")
}

type Man struct {
	person Person
}

func (m Man) eat() {
	print("child eat\n")
}

func main() {
	// https://blog.csdn.net/zhbinary/article/details/89418195
	// https://www.jianshu.com/p/fc7ba8876d00
	man := new(Man)
	man.eat()
	man.person.eat()
}
