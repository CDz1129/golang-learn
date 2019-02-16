package main

import (
	"com.cdz/learngo/basic/loop"
	"strings"
)

/**
 	- 总结for
		- 忽略初始条件,就相当于while
		- 死循环直接for {} 使用

	基本语法
	- for,if后条件语句不需要括号
	- if条件里也可以定义变量
	- 没有while
	- switch不需要break,也直接switch多个条件
*/
func main() {

	//forDemo()
	//
	//println(convert2Bin(5))
	//
	//println(convert2Bin(13))

	loop.PrintFile("basic/adc.txt")

	println("---------")
	s := `asd

	sss
	www
	111
	1`

	loop.PrintContentFile(strings.NewReader(s))
	//forever()
}
