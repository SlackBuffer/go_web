package main

import "fmt"

func main() {
	name := "Slack Buffer"
	tpl := `
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
	fmt.Println(tpl)
}
