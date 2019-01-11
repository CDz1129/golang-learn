package main

import (
	"fmt"
)

var temp = 1
var temp2 = "hello"

var (
	t = 1
	o = "hello"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //如果不打印的话,就会a与b变量名上会报错
}

func variableInitiaValue() {
	var a, b int = 1, 2           //可以看到,赋值时是可以一次赋多值
	var s string = "hello golang" // 在编译器中发现类型上画有黄线,提示go其实可以自动识别类型,它会根据我们的值去判断类型
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b = 1, 2
	var s = "hello golang" //自动判断类型,省略变量类型
	var c, d, z = 1, "hello", true
	fmt.Println(a, b, s)
	fmt.Println(c, d, z)
}

/**
变量定义简单写法
*/
func variableShorter() {

	a, b, c := 1, "hello", false //可以使用 := 的形式来定义变量(推荐使用)
	//a:=2	//报错 一个变量 只能有一次 := 的定义,所有语言都一样,变量定义在一个作用域中只能定义一次

	fmt.Println(a, b, c)
}

/**
产量的定义
*/
func constTest() {

	const a = 1
}

/**
函数外常量定义
*/
const name = 1

/**
枚举定义
*/
func enums() {

	/**
	通过 const括号集定义一个枚举
	*/
	const (
		up   = 1
		down = 2
	)

	/**
	其值可以使用 iota这个关键字来进行 下面的自增,
	如果想要跳过,可以使用_占位
	*/
	const (
		cpp = iota
		_
		java
		python
		golang
	)

	fmt.Println(cpp, java, python, golang)

	/**
	当然iota可以参与运算规则,下面所有都遵守这个规则
	*/
	// b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	//fmt.Println("hello world")
	//variableZeroValue()
	//variableInitiaValue()
	//variableTypeDeduction()
	//variableShorter()
	enums()
}
