package main

import (
	"com.cdz/learngo/retriever/mook"
	real2 "com.cdz/learngo/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

//定义一个post的接口
type Poster interface {
	Post(url string, params map[string]string) string
}

//接口组合： 定义一个 session接口
type Session interface {
	Poster
	Retriever
}

//定义一个session方法

func session(s Session) string {
	s.Post("http://www.imooc.com", map[string]string{"contents": "another a fake imooc.com"})
	s.Get("http://www.imooc.com")
	return "ok"
}

//定义一个方法，传入poster接口类型
func post(poster Poster) string {
	return poster.Post("http://www.imooc.com", map[string]string{"contents": "another a fake imooc.com"})
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	retriever := &mook.Retriever{"this is a fake imooc.com"}
	var r Retriever = retriever
	var p Poster = retriever
	download(retriever)
	post(retriever)
	session(retriever)

	fmt.Println(r)

	//s := download(r)

	//fmt.Println(s)
	fmt.Printf("%T %v\n", p, p)
	r = real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//interface是个什么东西？

	fmt.Println(r)
	//由此打印real.Retriever {Mozilla/5.0 1m0s}
	//interface 其内部有一个 interface 类型和值 (值也可以是指针)
	//fmt.Printf("%T %v\n",r,r)

	//fmt.Println(download(r))

	//如何判断 接口类型？

	//type assertion
	if retriever, ok := r.(real2.Retriever); ok {
		fmt.Println(retriever.UserAgent)
	} else if retriever, ok := r.(*mook.Retriever); ok {
		fmt.Println(retriever.Contents)
	} else {
		fmt.Println("类型错误")
	}

	//type switch
	switch v := r.(type) {
	case *mook.Retriever:
		fmt.Println("Contents:", v.Contents)
	case real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}

	/**
	查看接口变量：
		- 表示任何类型：interface{}
		- type assertion
		- type switch
	*/

}
