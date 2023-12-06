package control

import (
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
)

var basePath string

func init() {
	_, b, _, _ := runtime.Caller(0)
	basePath = filepath.Join(filepath.Dir(b), "..", "..")
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
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

func Router(pathWeb, tmpl, method string) {
	http.HandleFunc(pathWeb, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			RenderTemplate(w, tmpl)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
