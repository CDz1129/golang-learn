package main

import (
	"com.cdz/learngo/queue"
	"fmt"
)

func main() {

	queues := queue.Queue{1, 2}

	queues.Push(3)
	queues.Push(4)
	queues.Push(5)
	queues.Push("aasd")

	fmt.Println(queues.Pop())
	fmt.Println(queues.Pop())
	fmt.Println(queues.Pop())
	fmt.Println(queues.Pop())
	fmt.Println(queues.IsEmpty())
	fmt.Println(queues.Pop())
	fmt.Println(queues.Pop())
	fmt.Println(queues.IsEmpty())
}
