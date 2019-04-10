package main

import (
	"log"
	"os"
	"strings"
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

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var tpl *template.Template

func init() {
	// funcs have to be there before you parse
	// New's argument issue: https://golang.org/pkg/text/template/#Template.ParseFiles
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
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

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
