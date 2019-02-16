package main

import (
	"bufio"
	"com.cdz/learngo/functional/fib"
	"io"
)

func printContentFile(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		println(scanner.Text())
	}

}

func main() {

	fun := fib.Fibonacci()

	printContentFile(fun)
	//fmt.Println(fun())
	//fmt.Println(fun())
	//fmt.Println(fun())
	//fmt.Println(fun())
	//fmt.Println(fun())
	//fmt.Println(fun())
	//fmt.Println(fun())
	//fmt.Println(fun())
	//fmt.Println(fun())
	//fmt.Println(fun())
}
