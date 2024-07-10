// Templates Must

package main

import (
	"net/http"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"page.html",
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("page") // Criar um novo template
		t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
		t = template.Must(t.ParseFiles(templates...)) // Ajustar para a execução do template principal

		cursos := []Curso{
			{"Go", 40},
			{"Java", 30},
			{"Python", 50},
		}

		err := t.ExecuteTemplate(w, "page.html", cursos) // Ajustar para a execução do template principal
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
