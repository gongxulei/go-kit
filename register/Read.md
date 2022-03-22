## 服务注册中心



### 选项模式
> 使用选项模式来实现配置的设置，提高代码的可扩展性


```go
package main

import (
	"fmt"
)

type Options struct {
	Options1 string
	Options2 int
}
type OptionFun func(opts *Options)

func InitOptions(opts ...OptionFun) {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	fmt.Println("init options")
}

func WithStringOptions1(str string) OptionFun {
    return func(opts *Options) {
        opts.Options1 = str
	}
}



```
