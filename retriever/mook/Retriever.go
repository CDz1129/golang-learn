package mook

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever{Contents=%v}", r.Contents)
}

func (r *Retriever) Post(url string, params map[string]string) string {
	r.Contents = params["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
