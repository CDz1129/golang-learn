package main

import (
	"fmt"
	"sync"
)

/**
两种方法演示，不需要	time.Sleep(time.Millisecond)线程等待的方式，实现打印

方法二： GO自带程序实现 sync.WaitGroup{}


*/
func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 使用通信来共享内存，
		// 这里打印完了之后，给channel发送一条信息，说自己打印完了
		// 然后外面的程序接收，这样就可以实现，不用线程等待的方式，实现打印
		w.done()
	}
}

type worker struct {
	in   chan int
	done func() //函数式编程，结构体中 定义一个done的函数
}

func createWork(id int, wg *sync.WaitGroup) worker {

	w := worker{
		in: make(chan int),
		done: func() { //实现done这个函数的匿名函数
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {

	wg := sync.WaitGroup{}
	var channels [10]worker
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i, &wg)
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		channels[i].in <- 'a' + i
	}

	wg.Wait()

}

func main() {
	chanDemo()

}
