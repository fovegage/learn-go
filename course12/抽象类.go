package main

type AbstractPerson interface {
	//Sleep() func()
	Name() string
}

type Male struct {
}

func (m *Male) sleep(ap AbstractPerson) {
	//print(ap.Sleep())
	print("call:", ap.Name())
}

type Child struct {
	Male
}

func (c *Child) Name() string {
	return "child\n"
}

func main() {
	// http://neoyeelf.github.io/2019/04/07/golang%E4%B9%9F%E8%83%BD%E5%AE%9E%E7%8E%B0%E6%8A%BD%E8%B1%A1%E7%B1%BB%E4%BA%86%EF%BC%9F/

	// 类继承  后置接口 （r io.Reader）

	// 接口实现

	c := new(Child)
	// 向sleep传入实现类
	c.sleep(c)
}
