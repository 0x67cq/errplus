# 配置式 error 信息增幅工具
## 介绍
想要开发这个包的原因在于， 业务开发的时候，error过于简单，只能携带字符串，往往需要在不同的 代码位置的日志输出中加上不同的描述文字来进行错误发生位置的定位，或者在log包里加入某些行数定位， 但是报错堆栈信息始终是丢失的。往往有堆栈信息能够帮助开发者更好的定位问题，提高效率。

通过对1.13后的Error Wrap特性的利用。达到增加Error中携带栈,参数信息调试, 并且不影响errors.As等常用API的使用使用

支持可配置，保证调试环境和线上环境同一份代码，但是不影响线上性能

## Usage

尝试demo:
`export GOLANG_ERROR_STACK_DEBUG=GOLANG_ERROR_STACK_DEBUG && go run main.go`

作为import包使用:
```
package main

import (
	"errors"
	"fmt"

	"github.com/psychix/errstack/errstack"
)

func main() {
	fmt.Println("hello world")
	err := bllfunc()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func bllfunc() error {
	err := dotfunc(1, 2, 3)
	if err != nil {
		return errstack.Wrap(err, "name", 1, "stck", 2, "call", 3)
	}
	return nil
}

func dotfunc(...int) error {
	return errors.New("Ooooop!....")
}
```

