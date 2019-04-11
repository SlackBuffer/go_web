package main

import (
	"log"
	"os"
	"text/template"
	"time"
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
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	// 01 的位置就是月份，02 的位置就是日期
	// 01/02 03:04:05PM '06 -0700
	// Mon Jan 2 15:04:05 MST 2006
	return t.Format("01-02-2006")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
