package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node //指针
}

/**
工厂方法,控制创建
*/
func CreateNode(value int) *Node {
	return &Node{Value: value}
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

//此定义,方法接受者在前,这样的函数,其实和print(node Node)效果一样,
//只是调用形式上不同,
//func (node Node)Print() --> root.Print()
//func Print(node Node)   --> Print(root)
//那么思考,传入的是值还是指针?
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//因为之前我们学过,go语言中都是值传递,所以这里猜测也是
func (node *Node) SetValue(value int) {

	if node == nil {
		fmt.Println("nil调用了函数setValue")
		return
	}

	node.Value = value
	//发现结构的指针,使用不像内建类型一样麻烦,
	//直接使用变量名就好
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
