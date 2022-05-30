package main

import "fmt"

// Queue 基于slice
type Queue []int

func (q *Queue) Push(e int) {
	*q = append(*q, e)
}

func (q *Queue) Pop() int {
	fmt.Print("-----\n")
	fmt.Print(q, *q, (*q), "\n")
	fmt.Print("-----\n")
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
func (q *Queue) Print() {
	fmt.Println(*q)
}
func main() {
	q := Queue{1}
	q.Push(2)
	q.Print()
	q.Pop()
	q.Print()

}
