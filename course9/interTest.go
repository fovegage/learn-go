package main

type JPerson interface {
	eat() string
}

type Male struct {
	JPerson
}

func (m Male) eat() {
	print("eat\n")
}
func main() {
	male := new(Male)
	male.eat()
}
