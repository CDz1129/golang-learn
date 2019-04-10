package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testUserError string

func (u testUserError) Error() string {
	return u.Message()
}

func (u testUserError) Message() string {
	return string(u)
}

func userErrorHandler(writer http.ResponseWriter, request *http.Request) error {
	return testUserError("user error")
}

func notFindHandler(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func permissionHandler(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func okHandler(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprint(writer, "ok")
	return nil
}

var tests = []struct {
	hander  appHandler
	code    int
	message string
}{
	{errPanic, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)},
	{userErrorHandler, http.StatusBadRequest, "user error"},
	{notFindHandler, http.StatusNotFound, "Not Found"},
	{permissionHandler, http.StatusForbidden, "Forbidden"},
	{okHandler, http.StatusOK, "ok"},
}

func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		wrapper := errWrapper(tt.hander)
		wrapper(recorder, request)
		code := recorder.Code
		buffer := recorder.Body
		readAll, _ := ioutil.ReadAll(buffer)
		s := strings.Trim(string(readAll), "\n")
		if code != tt.code || s != tt.message {
			t.Errorf("输出（%d,%s）,应该得到（%d,%s）",
				recorder.Code, s,
				tt.code, tt.message)
		}

	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		h := errWrapper(tt.hander)
		//使用httptest.NewServer 直接起一个服务器测试
		server := httptest.NewServer(http.HandlerFunc(h))
		resp, _ := http.Get(server.URL)
		code := resp.StatusCode
		buffer := resp.Body
		readAll, _ := ioutil.ReadAll(buffer)
		s := strings.Trim(string(readAll), "\n")
		if code != tt.code || s != tt.message {
			t.Errorf("输出（%d,%s）,应该得到（%d,%s）",
				resp.StatusCode, s,
				tt.code, tt.message)
		}
	}
}
