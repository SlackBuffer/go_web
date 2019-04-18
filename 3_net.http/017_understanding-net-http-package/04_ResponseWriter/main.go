package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Ho-key", "this is from ho")
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(res, "<h1>hello</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
