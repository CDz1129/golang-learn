package main

import "fmt"

/**
函数学习
*/

/**
正常定义函数,(还是准寻golang定义原则变量名在前,变量类型在后,如果多个参数类型一致可以一起定义)
*/
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("错误的计算符号" + op)
	}
}

/**
多个返回值
	带余数除法
*/
func div(a, b int) (int, int) {
	return a / b, a % b
}

/**
函数式编程
	直接将函数做为一个参数传递
	这里就定义了一个func(a,b int) (r,q int)传入两个int值,返回两个int值的函数

注意这里只是做为演示,正常情况下返回多个参数会让人看起来很乱
一般多个参数返回,在有错误的情况下去做返回,返回错误信息
*/
func divFun(apply func(a, b int) (r, q int), a, b int) (r, q int) {
	return apply(a, b)
}

func eval1(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		r, _ := divFun(func(a, b int) (r, q int) {
			return a / b, a % b
		}, a, b) //由于golang中及其严格的语法限制,也就是定义的变量必须得用到,
		// 而这里我们调用得此函数去计算返回值有两个, 另一个没有用到
		//会编译不通过,解决办法,使用_使用下划线来代替

		return r, nil
	default:
		return 0, fmt.Errorf("错误的输入符号,%s", op)
	}
}

/**
在写函数得过程中,发现golang是没有函数重载,没有默认参数,可选参数这些东西
只有一个 就是可变参数 既...三个点代表
*/

func addAll(sum ...int) int {

	s := 0
	for e := range sum {
		s += sum[e]
	}

	return s
}

/**
最后是golang中指针学习
	golang中指针,并不像c++中指针,因为golang中指针是不能参与运算得
*/

/**
	指针学习,简单的值交接
 */
func swap(a, b int) {
	a, b = b, a
}

func main() {
	//println(eval(5, 5, "*"))
	//
	//i, i2 := div(11, 3)
	//
	//println(i,i2)
	//
	//r, q := divFun(func(a, b int) (r, q int) {
	//	return a / b, a % b
	//}, 11, 3)
	//
	//println(r,q)
	//
	//if i3, e := eval1(1, 10, ")");e!=nil{
	//	fmt.Println("Error:",e) //注意这里打印,不能使用 println()方法,会乱码,使用fmt.Println()不会
	//}else {
	//	println(i3)
	//}
	//
	//print(addAll(1, 2, 3, 4, 5))


	a,b:= 1,2

	swap(a,b)

	print(a,b)


	//var a = 1
	//var pa *int = &a
	//*pa = 3
	//println(a)
}
