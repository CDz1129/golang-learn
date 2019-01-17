package main

import "fmt"

func main() {

	//map定义
	//其他语言中,key必须是实现了hashcode 与 equals才可以作为key
	//首先map使用哈希表,必须可以比较相等
	//除了 slice map function 的内建类型都可以作为key
	//struct 类型不包含上诉字段 也可以作为key
	studens := map[string]string{
		"name": "张三",
		"age":  "18",
		"sex":  "男",
	}
	//多运行几遍,我们会发现 每次打印的map内部顺序是不一样的
	//说明key是无序的,如果需要有序打印,可以先将key全部取出,变成slice然后排序
	fmt.Println(studens)
	//获取key元素
	name := studens["name"]
	fmt.Println(name)
	//获取不存在元素
	idCard := studens["idCard"]
	fmt.Println(idCard) //打印空 ""所以我们怎么判断,key是否取到了呢?

	idCard,ok := studens["idCard"] //其实 获取map可以返回两个参数,第二个就是bool类型,false表示不存在
	fmt.Println(idCard,ok)
	if idCard,ok := studens["idCard"];ok {
		fmt.Println(idCard)
	}else {
		fmt.Println("key idCard not exist")
	}



	//创建 map
	var m2 map[string]int  //map == nil

	m3 := make(map[string]string) //map == empty map

	fmt.Println(studens,m2,m3)

	//遍历 map
	fmt.Println("----------key value一起遍历------\n")
	//key value一起遍历
	for k, v := range studens {
		fmt.Println(k,v)
	}
	fmt.Println("---------只遍历key-------\n")
	//只遍历key
	for k := range studens {
		fmt.Println(k)
	}

	fmt.Println("---------只遍历 value-------\n")
	//只遍历 value
	for _, v := range studens {
		fmt.Println(v)

	}

	//删除存在的键
	delete(studens, "age")
	fmt.Println("delete age after:",studens)

	//删除不存的键
	delete(studens,"idCard")
	fmt.Println("delete ont exist key after:",studens)
}
