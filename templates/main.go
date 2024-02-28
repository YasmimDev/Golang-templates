package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	curso := Curso{"Go", 40}
	temp := template.New("CursoTemplates")
	temp, _ = temp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")

	err := temp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
