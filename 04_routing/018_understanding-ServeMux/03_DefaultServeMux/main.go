package main

import (
	"io"
	"net/http"
)

/* type hotdog int
type hotcat int */

func main() {
	/* var d hotdog
	var c hotcat

	http.Handle("/dog/", d)
	http.Handle("/cat", c)
	*/
	http.HandleFunc("/dog/", d)
	// http.HandleFunc("/cat", c)
	http.Handle("/cat", http.HandlerFunc(c))

	// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux
	http.ListenAndServe(":8080", nil)
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy")
}
func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty")
}

/* func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy")
}
func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty")
} */
