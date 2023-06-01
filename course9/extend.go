package main

type smellable interface {
	smell()
}

type eatable interface {
	eat()
}

type fruit interface {
	smellable
	eatable
}

// 和下面的是等价的
//type fruit interface {
//	eat()
//	smell()
//}
