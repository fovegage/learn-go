package main

import "syscall/js"

// 编译生成wasm
// GOOS=js GOARCH=wasm go build -o static/main.wasm
// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static
func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World test!")
}
