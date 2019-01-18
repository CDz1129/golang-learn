package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {


	s := "Hi小智加油!"

	fmt.Println("len(s):",len(s)) //len(s): 15 为什么是15呢?

	for _, v := range []byte(s) {
		fmt.Printf("%X ",v)
		//48 69 E5 B0 8F E6 99 BA E5 8A A0 E6 B2 B9 21
		//可以看出从第三个开始(E5 B0 8F) -> 小 ,后面都是三个字节代表一个汉字
		//其实utf-8 是使用的可变长度来编码的(世界上这么多语言,如果全部都是3个直接来编码,太浪费空间)
	}

	fmt.Println()
	for k, v := range s { 	//v is a rune
		fmt.Printf("(%d,%X)",k,v)
	//(0,48)(1,69)(2,5C0F)(5,667A)(8,52A0)(11,6CB9)(14,21)
	//发现从第三个开始汉字,下标正确2,但是下一个就变成了5,
	//这说明,上面转为byte数组时我们的猜想是正确的,
	//一个汉字使用了3个byte来表示
	//但是这样遍历 根本不是我们想要的,因为下标的问题.

	//既 直接遍历字符串时,k 为 byte数组的下标 而 v 却是int32类型(rune)
	}
	fmt.Println()

	//使用 utf8工具类
	fmt.Println("s rune counts:",utf8.RuneCountInString(s))
	//s rune counts: 7 ,发现是s的长度是我们想要的了

	bytes := []byte(s)
	for len(bytes)>0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c " ,ch)
	}

	fmt.Println();
	for k, v := range []rune(s) {
		fmt.Printf("(%d %c)",k,v) //%c 可以将rune转为 字符串
	}

}
