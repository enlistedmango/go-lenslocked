package views

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil { // This error is if templates don't parse correctly
		return Template{}, fmt.Errorf("parsing template %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil { // This error is if we can't execute for whichever reason
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executig the template.", http.StatusInternalServerError)
		return
	}
}
