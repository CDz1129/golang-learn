package main

import "fmt"

/**
查找到一个字符串中
最长不重复字符串

例如：
asdasdss -> asd
*/

func findMaxNoRepeatString(s string) int {

	start := 0
	keysIndex := make(map[rune]int)
	lenth := 0

	for i, v := range []rune(s) {
		lastIndex, ok := keysIndex[v]
		if ok && lastIndex >= start {
			start = lastIndex + 1
		} else {
			lenth = i - start + 1
		}
		keysIndex[v] = i
	}
	return lenth

}

func main() {
	fmt.Println(findMaxNoRepeatString("abc"))
	fmt.Println(findMaxNoRepeatString("abccba"))
	fmt.Println(findMaxNoRepeatString("aaaaa"))
	fmt.Println(findMaxNoRepeatString(""))
	fmt.Println(findMaxNoRepeatString("asdfghjk"))
	fmt.Println(findMaxNoRepeatString("我是小智"))
	fmt.Println(findMaxNoRepeatString("小智智小"))
}
