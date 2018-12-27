package main

import (
	"io/ioutil"
	"fmt"
)

/**
	if语言demo,可以看出 if语句后条件不需要加()
 */
func ifDemo() {

	const filename = "adc.txt"

	bytes, e := ioutil.ReadFile(filename) //golang返回值可以是多个,需要多个值来接收

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%s\n", bytes)
	}

	if file, i := ioutil.ReadFile(filename); i != nil { //另一种写法,if语句中是可以跟多个语句,必须使用;分号隔开,条件中可以赋值
		fmt.Println(e)
	} else {
		fmt.Printf("%s/n", file)
	}

}

/**
	switch用法,不需要添加 break,会自动返回
	switch后可以没有表达式
 */
func switchDemo() {
	fmt.Print(grade(81),grade(20),grade(100),grade(90))
}

func grade(soure int) string {

	g:=""
	switch {
	case soure<60:
		g="F"
	case soure<80:
		g="D"
	case soure<90:
		g = "B"
	case soure<=100:
		g="A"
	case soure<0||soure>100:
		panic(fmt.Sprintf("Wrong score: %d",soure))
	}
	return g
}

func main() {

	//ifDemo()

	switchDemo()
}
