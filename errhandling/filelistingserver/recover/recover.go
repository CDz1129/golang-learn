package main

import (
	"errors"
	"log"
)

func tryRecover() {
	defer func() {
		r := recover()
		if e, ok := r.(error); ok {
			log.Println(e)
		}
	}()
	panic(errors.New("recover错误处理演示"))
}

func main() {
	tryRecover()
}
