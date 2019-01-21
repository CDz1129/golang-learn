package main

import (
	"fmt"
	"golang-learn/tree"
)

func main() {

	//创建 和 调用
	var root tree.Node

	root.Value = 1
	root.left = &treeNode{value: 5}
	root.right = new(treeNode) //new 返回的是指针
	root.left.right = createNode(0)
	root.right.left = &treeNode{value: 10, left: nil}
	root.right.left.right = &treeNode{
		value: 15,
		left: &treeNode{
			value: 0,
			left:  nil,
		}}

	//root.right.left.right.right.setValue(1)
	/**
	nil调用了函数setValue //看到打印了 这句话 说明是被调用了
	但是下面还是报错了,这是因为 nil 没有value这个值才报的错,
	解决办法直接return
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xc0000005 code=0x1 addr=0x0 pc=0x48b9a3]

	goroutine 1 [running]:
	main.(*treeNode).setValue(0x0, 0x1)
	*/
	//方法编写
	root.print()

	root.right.print()
	root.right.setValue(5)
	//func (node treeNode)setValue(value int)
	root.right.print()
	//{0 0xc04204c440 <nil>}
	//{0 0xc04204c440 <nil>}
	//发现两次打印一样,并没有修改值,
	//所以也是值传递

	//如何修改值?
	//传入指针
	//func (node *treeNode)setValue(value int)

	/**
	- 创建方式很多,没有构造函数
		- 使用工厂方法,实现类似构造函数
	- 调用结构体内部内容,直接使用.点
		- 无论指针还是值,都使用.点来调用
		- nil也可以调用方法
	- 方法不特殊声明的话,是值传递
	- 修改值必须使用引用传递
	- 无论方法接受者是 值 还是 指针,调用者可以使用 指针 也可以使用值 调用(编译器自己判断)
	*/

	fmt.Println("--------------traverse----------")

	root = treeNode{value: 5}
	root.left = &treeNode{value: 1}
	root.right = &treeNode{value: 2}
	root.left.right = &treeNode{value: 3}
	root.right.left = &treeNode{value: 4}

	root.traverse()
}
