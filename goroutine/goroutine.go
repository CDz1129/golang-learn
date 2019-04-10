package main

import (
	"fmt"
	"time"
)

func main() {
	var a [1000]int
	for i := 0; i < 1000; i++ {
		//携程是非抢占式的，但是发现打印出来的结果
		//还是像多线程切换一样打印，（非抢占式，是指携程主动交出控制权）
		//很明显我们代码中没有写，主动交出控制权的代码
		//其实主要是**println** io会自动交出控制权
		go func(i int) { //race condition
			for {
				//**println** io会自动交出控制权
				fmt.Println(i)
				//使用数组，并累加数组所在位置的数字，最后打印数组
				//看是否有切换操作
				//a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
