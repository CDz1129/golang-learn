package main

import "fmt"

func main() {
	//arr := [...]int{0,1,2,3,4,5,6,7}
	//s1 := arr[2:6]// [3 4 5 6]左闭右开,一般计算机语言中都是这么定义
	////slice不是值类型,底层是对其数组的封装
	//fmt.Println("arr[2:6]:",s1)
	//s2 := arr[:6]
	//fmt.Println("arr[:6]:", s2)
	//s3 := arr[2:]
	//fmt.Println("arr[2:]:", s3)
	//s4 := arr[:]
	//fmt.Println("arr[:]:", s4)
	//fmt.Println("s1:",s1)
	//fmt.Println("------------updateSlice before----------")
	//updateSlice(s1)
	//fmt.Println("------------updateSlice after----------")
	//fmt.Println("s1:",s1)
	//fmt.Println("s2:",s2)
	//fmt.Println("s3:",s3)
	//fmt.Println("s4:",s4)
	//fmt.Println("arr:",arr)
	/**
	 - slice不是值类型,go语言中除了slice所有都是值类型(slice的特殊之处就提现在这个地方)
		- slice在内部,是有一个数据结构的,是对arr的一个view(视图)

	slice数据结构:

	package runtime

	import (
		"unsafe"
	)

	type slice struct { //slice的底层数据结构
		array unsafe.Pointer
		len   int
		cap   int
	}
	 */

	//s4 = s4[:6]
	//fmt.Println("reslice s4:",s4)
	//s4 = s4[2:]
	//fmt.Println("reslice s4:",s4)

	//拓展
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1,s2:",s1,s2)
	fmt.Println("s1[4]",s1[4])

}
func updateSlice(ints []int) {
	ints[2] = 100
}
