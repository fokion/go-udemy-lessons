package main

import (
	"html/template"
	"os"
	"log"
)

var tmpl *template.Template
func init(){
	tmpl = template.Must(template.ParseGlob("views/*.html"))
}
func main()  {
	err := tmpl.ExecuteTemplate(os.Stdout,"index.html","Hello World")
	if(err !=nil){
		log.Fatal(err)
	}
}
