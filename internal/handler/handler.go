package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	renderTemplate("view/index.html", w)
}

func HandleDesc(w http.ResponseWriter, r *http.Request) {
	renderTemplate("view/desc.html", w)
}

func HandleKeyboards(w http.ResponseWriter, r *http.Request) {
	renderTemplate("view/keyboards.html", w)
}

func renderTemplate(tmpl string, w http.ResponseWriter) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Println("skill issue", err)
	}
	t.Execute(w, "")
}
