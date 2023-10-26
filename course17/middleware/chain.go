package middleware

import (
	"fmt"
	"math"
	"net/http"
)

const maxIndex = math.MaxInt8 / 2

type myFuncHandle func(ctx *MyContext)

type MyContext struct {
	handleChains []myFuncHandle
	Req          *http.Request
	index        int8
}

func (m *MyContext) Next() {
	if m.index < maxIndex {
		m.index++
		m.handleChains[m.index](m)
	}
}

func (m *MyContext) Abort() {
	m.index = maxIndex
	fmt.Println("已被终止...")
}

func (m *MyContext) Use(f myFuncHandle) {
	m.handleChains = append(m.handleChains, f)
}

func (m *MyContext) Get(uri string, f myFuncHandle) {
	m.handleChains = append(m.handleChains, f)
}

func (m *MyContext) Run() {
	m.handleChains[0](m)
}
