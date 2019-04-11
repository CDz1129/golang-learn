package main

import (
	"fmt"
	"sync"
	"time"
)

/**
通过 Mutex实现一个atomitInt
*/

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *AtomicInt) increment() {
	fmt.Println("safe increment")
	func() { //想要一块代码区，安全
		//使用 func方法包装起来
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *AtomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {

	var a AtomicInt

	a.increment()

	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())

}
