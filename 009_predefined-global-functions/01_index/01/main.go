package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xs := []string{"a", "b", "c", "d"}
	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
