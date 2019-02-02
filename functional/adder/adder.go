package main

import "fmt"

func Adder() func(int) int {
	num := 0
	return func(i int) int {
		num += i
		return num
	}
}

func main() {
	adder := Adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+...+%d=%d\n", i, adder(i))
	}
}
