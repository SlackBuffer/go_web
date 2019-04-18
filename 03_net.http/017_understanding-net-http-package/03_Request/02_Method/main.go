package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// has to call `r.ParseForm` first in order to use `r.Form`
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data) // for values of both form and url
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
