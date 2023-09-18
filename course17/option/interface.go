package option

// ServerOption 定义函数选项
type ServerOption struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type IOption interface {
	apply(option *ServerOption)
}

// FuncOption 接收函数选项
type FuncOption struct {
	f func(option *ServerOption)
}

func (fo FuncOption) apply(o *ServerOption) {
	fo.f(o)
}

func NewFuncOption(f func(o *ServerOption)) IOption {
	return FuncOption{
		f: f,
	}
}

func WithHost(host string) IOption {
	return NewFuncOption(func(o *ServerOption) {
		o.Host = host
	})
}

func BuildServer(opts ...IOption) *ServerOption {
	server := &ServerOption{
		Port: 9999,
	}
	for _, o := range opts {
		o.apply(server)
	}
	return server
}
