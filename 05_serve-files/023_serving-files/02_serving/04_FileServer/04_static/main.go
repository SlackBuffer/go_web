package main

import (
	"log"
	"net/http"
)

func main() {
	// 目录下有 index.html 则直接展示 index.html，不展示目录内容
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
