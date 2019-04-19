package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", cat)
	http.HandleFunc("/cat.jpeg", catPic)
	http.ListenAndServe(":8080", nil)
}

func cat(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<img src="cat.jpeg" />
	`)
}

func catPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("cat.jpeg")
	if err != nil {
		http.Error(w, "file not fount", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
