package main

import (
	"com.cdz/learngo/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(
	handler appHandler) func(
	writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//错误处理函数也需要recover
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		e := handler(writer, request)
		if e != nil {
			log.Printf("Error occurred handling request:%s", e)
			if userErr, ok := e.(UserError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(e):
				code = http.StatusNotFound
			case os.IsPermission(e):
				code = http.StatusForbidden //没有权限
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)

		}

	}
}

type UserError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.Handerlist))
	e := http.ListenAndServe(":8888", nil)
	if e != nil {
		panic(e)
	}
}
