package main

import "fmt"

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
	for k, v := range s {
		fmt.Printf("(%d,%X)",k,v)
	}
	fmt.Println()

}
