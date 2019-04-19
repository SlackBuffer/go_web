package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", cat)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func cat(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Fprintln(w, "go look at the terminal")
}
