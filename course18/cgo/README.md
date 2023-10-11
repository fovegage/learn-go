## cgo测试

```
# 交叉编译
https://www.topgoer.com/%E5%85%B6%E4%BB%96/%E8%B7%A8%E5%B9%B3%E5%8F%B0%E4%BA%A4%E5%8F%89%E7%BC%96%E8%AF%91.html

# 注意运行要指定cgo(CGO_ENABLED=1)
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build cgo_exp1.go (linux)
GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go run cgo_exp1.go

# 参考文档（https://juejin.cn/post/7047405294107754533）
# 生成动态链接库   macos dylib
gcc -arch arm64 -arch x86_64 -fPIC -shared -o lib/libadd.dylib include/add.c

# 生成二进制文件
gcc include/test.c -L lib/ -ladd -o test


# 查看依赖
nm -D libadd.dylib
otool -L libadd.dylib

# cgo
https://www.ithere.net/article/1506646292826779650
```

### 问题

```
# building for macOS-x86_64 but attempting to link with file built for macOS-arm64
https://stackoverflow.com/questions/74797526/building-for-macos-x86-64-but-attempting-to-link-with-file-built-for-macos-arm64
```