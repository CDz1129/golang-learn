package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	生成一个channel
并且	数据发送 时间随机，
	数据数据 自增
*/
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out

}

func work(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}

}

func createWork(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}

func main() {

	//var v1, v2 chan int //v1 and v2 is nil

	v1, v2 := generator(), generator()
	work := createWork(0)

	after := time.After(10 * time.Second)
	tick := time.Tick(1 * time.Second)
	var values []int
	for {
		var activeWorker chan<- int //使用到select 中 如果 channel是nil的话 不case
		var activeValue int
		if len(values) > 0 {
			activeWorker = work
			activeValue = values[0]
		}
		select {
		case n := <-v1:
			values = append(values, n)
		case n := <-v2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]

		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("value queue len:", len(values))
		case <-after:
			fmt.Println("bye bye")
			return
		}
	}
}
