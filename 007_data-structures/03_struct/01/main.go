package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	a := sage{
		Name:  "sb",
		Motto: "do",
	}

	err := tpl.Execute(os.Stdout, a)
	if err != nil {
		log.Fatalln(err)
	}
}
