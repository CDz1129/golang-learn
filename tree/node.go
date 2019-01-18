package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode //指针
}

/**
工厂方法,控制创建
 */
func createNode(value int) *treeNode {
	return &treeNode{value: value}
	/**
	注意 :
			这里因为我一直做Java觉得这样写没有什么问题,
			而其实这里是 返回的局部变量的地址,
			在Java中没有问题,是因为Java中几乎所有的对象都存放在堆中
			这段代码在c++中其实是有问题的,
			也是不同语言之间的区别,地址在什么地方?
			我理解的是,c++ 与 go 地址可以是 在栈上创建,
			所以在c++中是会有问题的,因为一旦调用完栈就会弹出
			而go中一个地址 到底是在堆还是栈上?
			不用管! 编译器会智能判断,如果后续需要用到,地址就会创建在堆上
					如果用不到就会创建在栈上
	 */

}

//此定义,方法接受者在前,这样的函数,其实和print(node treeNode)效果一样,
//只是调用形式上不同,
//func (node treeNode)print() --> root.print()
//func print(node treeNode)   --> print(root)
//那么思考,传入的是值还是指针?
func (node treeNode) print() {
	fmt.Print(node.value," ")
}

//因为之前我们学过,go语言中都是值传递,所以这里猜测也是
func (node *treeNode) setValue(value int) {

	if node == nil {
		fmt.Println("nil调用了函数setValue")
		return
	}

	node.value = value
	//发现结构的指针,使用不像内建类型一样麻烦,
	//直接使用变量名就好
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {

	//创建 和 调用
	var root treeNode
	root.value = 1
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


	 root = treeNode{value:5}
	 root.left = &treeNode{value:1}
	 root.right = &treeNode{value:2}
	 root.left.right = &treeNode{value:3}
	 root.right.left = &treeNode{value:4}

	 root.traverse()
}
