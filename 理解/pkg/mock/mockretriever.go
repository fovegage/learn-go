package mock

// 接口的定义   定义了接口就要实现它
type RetrieverHandler interface {
	Get(url string) string
}

// 定义基础鸭子结构
type Retriever struct {
	Content string
}

// 具体的实现
// 这两块是相互依赖的
// 下面方法的依赖
func (r Retriever) Get(url string) string {
	return url
}
