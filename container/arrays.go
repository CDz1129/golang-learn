package main

import "fmt"

/**
array数组基本操作
 */

/**
创建一个数组
 */
func createArrays() {

	//第一种 定义数组
	var arr1 [5]int

	//第二种 定义数组方式
	arr2 := [3]int{1, 2, 3}

	//可变长度的数组
	arr3 := [...]int{1, 2, 3, 4, 5}

	//二维数组
	var arr4 [4][5]int

	//原始fori的形式
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//遍历 之遍历出下标 使用range关键字
	for e := range arr3 {
		//e 下标
		fmt.Println(e, arr3[e])
	}

	//关键字range遍历出 下标与值
	for k, v := range arr3 {
		//k下标 v 值
		fmt.Println(k, v)
	}

	fmt.Println(arr1, arr2, arr3, arr4)

	printArr(arr3)
	fmt.Println("---------------update before--------------")
	//updateArr(arr3)
	updateArrByPointer(&arr3)
	fmt.Println("---------------update after--------------")
	printArr(arr3)

}
func updateArrByPointer(ints *[5]int) {
	ints[0] = 1000
	fmt.Println("-------------update success----------")
}
func updateArr(ints [5]int) {
	ints[0] = 1000
	fmt.Println("-------------update success----------")
}

func printArr(ints [5]int) { //使用数组作为参数,就需要指定数量 ,
	//[5]int类型如果传入[3]int就会报错.
	//如果不想被数量限制可以使用 [...]int作为参数类型
	for k, v := range ints {
		fmt.Println(k, v)
	}
}

func main() {
	createArrays()
}
