package main

import (
	"fmt"
)

/**
两种方法演示，不需要	time.Sleep(time.Millisecond)线程等待的方式，实现打印

第一种：自己实现 对外通信

*/
func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 使用通信来共享内存，
		// 这里打印完了之后，给channel发送一条信息，说自己打印完了
		// 然后外面的程序接收，这样就可以实现，不用线程等待的方式，实现打印
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWork(id int) worker {

	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {

	var channels [10]worker
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		channels[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		<-channels[i].done
	}

}

func main() {
	chanDemo()

}
