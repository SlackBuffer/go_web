package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// method 为 get 参数会在 url 中出现
	// method 为 post 参数在 request body 里
	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form><br>
	`+v)
}
