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
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s1 := arr[2:6]
	//s2 := s1[3:5]
	//fmt.Println("arr:",arr)
	//fmt.Printf("s1 v:%v, len(s1):%d, cap(s1):%d\n",s1,len(s1),cap(s1))
	//fmt.Printf("s2 v:%v, len(s2):%d, cap(s2):%d\n",s2,len(s2),cap(s2))

	//slice 添加 复制 删除
	//添加
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]

	ss1 := append(s1, 10)
	ss2 := append(ss1, 10)
	ss3 := append(ss2, 10)

	fmt.Println("arr:", arr)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("ss1:", ss1)
	fmt.Println("ss2:", ss2)

	/**
	arr: [0 1 2 3 4 5 10 10]
	s1: [2 3 4 5]
	s2: [5 10]
	ss1: [2 3 4 5 10] ss1 是对s1的追加,我们发现其追加完后会再次返回一个slice
							且看到arr这个数组的后面两个值被改变了 而arr数组改变,s2也会就会跟着变化
	ss2: [2 3 4 5 10 10]
	 */
	fmt.Println("ss3:", ss3)
	fmt.Printf("ss3 v:%v, len(ss3):%d, cap(ss3):%d\n", ss3, len(ss3), cap(ss3))
	fmt.Printf("arr v:%v, len(arr):%d, cap(arr):%d\n", arr, len(arr), cap(arr))
	/**
	arr: [0 1 2 3 4 5 10 10]
	s1: [2 3 4 5]
	s2: [5 10]
	ss1: [2 3 4 5 10]
	ss2: [2 3 4 5 10 10]
	ss3: [2 3 4 5 10 10 10] 	当我们继续添加一个元素,因为此时已经是超过arr的长度
								此时,打印发现arr不再变化,而返回的ss3是追加上的
	ss3 v:[2 3 4 5 10 10 10], len(ss3):7, cap(ss3):12
								打印出ss3的信息可以看出,其实已经是一个行的slice了,跟arr没有任何关系
								疑问 为什么 cap为12?
								要回答为什么是12我们得再做一个实验
	 */

	/**
	解答cap得长度秘密
	 */
	var s []int
	for i := 0; i < 10; i++ {
		fmt.Printf("s v:%v, len(s):%d, cap(s):%d\n", s, len(s), cap(s))
		s = append(s, i)
	}
	/**
	s v:[], len(s):0, cap(s):0
	s v:[0], len(s):1, cap(s):1
	s v:[0 1], len(s):2, cap(s):2
	s v:[0 1 2], len(s):3, cap(s):4
	s v:[0 1 2 3], len(s):4, cap(s):4
	s v:[0 1 2 3 4], len(s):5, cap(s):8
	s v:[0 1 2 3 4 5], len(s):6, cap(s):8
	s v:[0 1 2 3 4 5 6], len(s):7, cap(s):8
	s v:[0 1 2 3 4 5 6 7], len(s):8, cap(s):8
	s v:[0 1 2 3 4 5 6 7 8], len(s):9, cap(s):16

	打印结果可以发现,其实cap是有规律得变化得.
	初始时我们创建了一个空的slice s ,再golang中没有null得概念,有nil(nil会初始化数据),但是nil是可以调用复制得
	开始0 后面1 超过1cap后变成2 后面是4,由此可以看出,cap的长度是,当cap长度不足时,扩充2倍
	 */

	//create slice
	s1 = arr[:]
	s1 = make([]int, 12)     //第一个参数指定len长度
	s1 = make([]int, 12, 16) //第二个参数可以指定 cap

	//copy slice
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[:]
	copyArr := make([]int, 12, 16)
	copy(copyArr, s1) //copy参数,第一个是被赋值的对象,第二个是数据源
	fmt.Println(copyArr)

	//将copyArr 中第三个元素删掉
	copyArr = append(copyArr[:3], copyArr[4:]...)
	//删除其中元素是没有函数的
	//所以我们可以间接的实现,截取前半部分与后半部分,然后再合起来
	//我们看到append 第二个参数slice后面有...三个点,
	//这是因为在append函数第二个参数不能传入slice但是可以加上三个点就可以了
	fmt.Println(copyArr)

	//取出头元素
	//popping from front

	fmt.Println("copyArr", copyArr)
	front := copyArr[0]
	copyArr = copyArr[1:]
	fmt.Println("front:", front)
	fmt.Println("after popping from front, copyArr:", copyArr)

	//取出最后一个元素
	//poping from back
	tail := copyArr[len(copyArr)-1]
	copyArr = copyArr[:len(copyArr)-1]

	fmt.Println("tail:" ,tail)
	fmt.Println("after poping from back ,copyArr:",copyArr)

}

func updateSlice(ints []int) {
	ints[2] = 100
}
