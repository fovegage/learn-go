- 切片包含：底层数组的指针、切片的长度和切片的容量

- Go语言和Python的切片是不同的，Go语言不支持切片步长和负索引

- make([]int, len, cap):若len或cap只提供一个,则len=cap