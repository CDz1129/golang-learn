package fib

import (
	"fmt"
	"io"
	"strings"
)

/**
函数生成斐波那契数列（每个数字都是由前两个数字相加得到）

*/
func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		//在这里，生成关键
		// 1 1 2 3 5 8
		//   a b
		//	   a b
		a, b = b, a+b
		return a
	}
}

//函数实现接口
//定义一个函数，使用type修饰。可以实现接口，也就是说只要是被type修饰的东西都可以实现接口。
type IntGen func() int

//实现read接口
func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()

	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
