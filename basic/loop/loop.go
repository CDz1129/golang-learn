package loop

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

/**
正常的for循环使用

- for条件不需要括号
- for可以省略初始条件,递增条件,判断条件
*/
func forDemo() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

/**
	- for可以没有初始条件

转换二进制
	一个(n)十进制的数转换为2进制 步骤
			n / 2 得到的数字 % 2 所得到的数字 向前放置
			当 n / 2 为0时结束
*/
func convert2Bin(n int) string {

	bin := ""
	for ; n > 0; n /= 2 {
		bin = strconv.Itoa(n%2) + bin
	}
	return bin
}

/**
读取文件时
	使用os.open(string文件名)打开一个文件
	然后通过bufio.NewScanner(file) 生成一个文件的扫描
	for只有一个判断语言,就能类似与 while循环语言 ,
	值得一提的是golang中没有while语句
*/
func PrintFile(filename string) {

	file, e := os.Open(filename)

	if e != nil {
		panic(e)
	}

	PrintContentFile(file)

}

func PrintContentFile(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		println(scanner.Text())
	}

}

/**
对于死循环来说,
	golang有更加简洁的写法
	直接使用for不加条件,就是死循环
	相比java -> 死循环 for(;;) 或者用的最多的 while(true)  会更加简洁
*/
func forever() {
	for {
		println("a")
	}
}
