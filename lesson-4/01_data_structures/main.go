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
	names := []string{"Fokion","Alexandra","Natalia"}
	err := tmpl.ExecuteTemplate(os.Stdout,"index.html",names)
	if(err !=nil){
		log.Fatal(err)
	}
}
