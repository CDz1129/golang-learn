package main

import (
	"bufio"
	"com.cdz/learngo/functional/fib"
	"fmt"
	"os"
)

func tryDefer() {
	//file, err := os.Open("README.md")
	//if err!=nil {
	//	panic(err)
	//}
	//defer file.Close()
	//loop.PrintContentFile(file)
	//defer fmt.Println(2)
	//fmt.Println(1)
	//defer fmt.Println(3)
	//return
	/**
	打印：
	1
	3
	2
	说明在defer中，只有在方法执行完前，或者方法return之前才会执行
	不影响原有代码，
	且多个defer是有一个栈的，遵循后进先出的原则。
	*/

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 10 {
			panic("num too big")

		}
	}
}

func writFile(filename string) {

	//file, e :=os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)

	file, e := os.Create(filename)
	//e = errors.New("this is a custom error")
	if e != nil {
		if pathError, ok := e.(*os.PathError); !ok {
			panic(e)
		} else {
			fmt.Printf("%s, %s ,%s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	defer writer.Flush()
	fibonacci := fib.Fibonacci()
	for i := 0; i <= 20; i++ {
		fmt.Fprintln(writer, fibonacci())
	}

}

func main() {
	//tryDefer()

	writFile("fib.txt")
}
