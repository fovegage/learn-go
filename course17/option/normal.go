package option

type Option func(*Options)

type Options struct {
	Name string
	Age  int
}

func WithAge(age int) Option {
	// 外部传入修改
	return func(o *Options) {
		// 这儿相当于更新了Options 将传入的重新复制对age
		println(o.Name, o.Age)
		// 刷新一下值
		o.Age = age
	}
}

func NewOptions(opts ...Option) *Options {
	o := &Options{
		Name: "gage",
	}
	for _, opt := range opts {
		opt(o)
	}
	// 值传递
	println("modify:", o.Age)
	return o
}
