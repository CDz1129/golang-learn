package main

import (
	"fmt"

	"com.cdz/learngo/tree"
)

func main() {

	//创建 和 调用
	var root tree.Node

	root.Value = 1
	root.Left = &tree.Node{Value: 5}
	root.Right = new(tree.Node) //new 返回的是指针
	root.Left.Right= tree.CreateNode(0)
	root.Right.Left = &tree.Node{Value: 10, Left: nil}
	root.Right.Left.Right= &tree.Node{
		Value: 15,
		Left: &tree.Node{
			Value: 0,
			Left:  nil,
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
	root.Print()

	root.Right.Print()
	root.Right.SetValue(5)
	//func (node treeNode)setValue(value int)
	root.Right.Print()
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

	root = tree.Node{Value: 5}
	root.Left = &tree.Node{Value: 1}
	root.Right = &tree.Node{Value: 2}
	root.Left.Right= &tree.Node{Value: 3}
	root.Right.Left = &tree.Node{Value: 4}

	root.Traverse()
}
