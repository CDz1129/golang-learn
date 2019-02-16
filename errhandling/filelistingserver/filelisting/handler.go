package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix string = "/list/"

func Handerlist(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start " + prefix)
	}

	path := request.URL.Path[len(prefix):]
	file, e := os.Open(path)
	if e != nil {
		return e
	}
	defer file.Close()
	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		return e
	}
	writer.Write(bytes)
	return nil
}

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}
