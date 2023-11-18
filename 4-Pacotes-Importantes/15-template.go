package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	templates := []string{
		"public/content.html",
		"public/header.html",
		"public/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").Funcs(template.FuncMap{"ToUpper": strings.ToUpper}).ParseFiles(templates...))
		err := t.Execute(w, []Curso{
			{"Go", 40},
			{"Python", 40},
			{"JavaScript", 40},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8090", nil)
}
