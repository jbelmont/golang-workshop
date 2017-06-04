package template

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name   string
	Age    int
	Weight int
	City   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("users.html"))
}

func templates() string {
	p := person{
		Name:   "John Rambo",
		Age:    37,
		Weight: 205,
		City:   "Ft. Bragg",
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
	return p.Name
}
