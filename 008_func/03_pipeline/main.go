package main

import (
	"log"
	"math"
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

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var tpl *template.Template

func init() {
	// funcs have to be there before you parse
	// New's argument issue: https://golang.org/pkg/text/template/#Template.ParseFiles
	// demo https://play.golang.org/p/XnpQkzL1ekQ
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
