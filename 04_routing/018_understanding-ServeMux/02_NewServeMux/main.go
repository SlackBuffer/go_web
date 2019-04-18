package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	// 会捕获子路由
	mux.Handle("/dog/", d)

	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy")
}
func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty")
}
