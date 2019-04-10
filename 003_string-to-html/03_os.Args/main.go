package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	str := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello, 世界</title>
	</head>
	<body>
	<h1>
	` + name + `</h1>
	</body>
	</html>
	`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
