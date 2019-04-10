package main

import (
	"fmt"
	"time"
)

func work(id int, c chan int) {
	//for {
	//	//如果有关闭channel的情况，当关闭之后，还是会接收到，只是接收到的是channel类型的初始值 int为0 string为''
	//	//解决：
	//	//一 接收时 接收两个值 第二个为bool判断是否发送完
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}

	//二 使用range循环接收，没有了会自动结束
	//此方法同样适用于没有关闭的channel，当没有关闭 就会一直循环下去
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}

}

/**
此时我们，返回的这个channel其实还是有点问题，不能做限制
这里业务简单，可以看出 返回的channel只能发送不能接受，当业务复杂就不那么容易看出
所以，可以指定返回channel的类型
	<- chan 说明是接受类型
	chan <- 说明是发送类型
*/
func createWork(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}

func chanDemo() {
	//var c chan int 此时只是创建一个channel为nil 目前是不能用的
	//channel为一等公民 可以作为参数 和返回值
	//c := make(chan int)
	//go work(0,c)
	//c <- 1//发数据
	//c <- 2

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferChannel() {
	c := make(chan int, 3)
	/**
	make多传入一个参数，说明是带有缓冲区的channel,如果没有缓冲区，对于创建的channel
	只要发送了，就一定要接收，还是会有一定的切换成本

	当多余缓冲区发送，同样也会报错deadlock!
	*/
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
}

func closeChannel() {
	c := make(chan int)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
	//bufferChannel()
	//closeChannel()

}
