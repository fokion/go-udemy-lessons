package main

import (
	"html/template"
	"os"
)

var tmpl *template.Template
func init(){
	tmpl = template.Must(template.ParseGlob("views/*.html"))
}

type Person struct {
	Name string
	Surname string
}
func main()  {
	var people = []Person{}
	p1 := Person{"Fokion","Sotiropoulos"}
	people = append(people,p1)
	people = append(people,Person{"Natalia", "Sotiropoulou"})
	tmpl.ExecuteTemplate(os.Stdout,"index.html",p1)
	tmpl.ExecuteTemplate(os.Stdout,"index-people.html",people)
}
