.PHONY: all
all: build cp

build:
	GOOS=js GOARCH=wasm go build -o static/main.wasm

cp:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static
