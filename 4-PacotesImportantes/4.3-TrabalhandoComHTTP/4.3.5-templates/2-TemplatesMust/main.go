// Templates Must

package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{
		Nome:         "Golang",
		CargaHoraria: 40,
	}

	t := template.Must(template.New("curso").Parse("Curso: {{.Nome}}\nCarga Horária: {{.CargaHoraria}} horas\n"))

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
