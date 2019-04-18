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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
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
	b := sage{
		Name:  "sb",
		Motto: "it",
	}

	sages := []sage{a, b}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}
	t := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}
	cars := []car{f, t}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
