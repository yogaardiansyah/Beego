package control

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, basePath, tmpl string) {
	templatePath := filepath.Join(basePath, "views", tmpl)
	temp, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Router(basePath, tmpl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, basePath, tmpl)
	}
}
