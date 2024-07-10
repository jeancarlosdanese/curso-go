// Templates

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

	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Bem-vindo ao curso de {{.Nome}} com carga horária de {{.CargaHoraria}} horas.")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
