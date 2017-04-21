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
	p1 := Person{"Fokion","Sotiropoulos"}
	tmpl.ExecuteTemplate(os.Stdout,"index.html",p1)
}
