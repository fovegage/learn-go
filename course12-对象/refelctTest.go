package main

import "reflect"

type RPerson struct {
}

func (r *RPerson) Eat() {
	print("test eat")
}
func main() {
	rp := new(RPerson)
	value := reflect.ValueOf(rp)
	caller := value.MethodByName("Eat")
	caller.Type().Out(1)
}
