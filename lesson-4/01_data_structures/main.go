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
	namesInCities := map[string]string{
		"Edi":"Fokion",
		"Gla":"Natalia",
	}
	err := tmpl.ExecuteTemplate(os.Stdout,"index.html",names)
	if(err !=nil){
		log.Fatal(err)
	}
	fl,err := os.Create("out.html")
	if(err !=nil){
		log.Fatal(err)
	}
	defer fl.Close()
	err = tmpl.ExecuteTemplate(fl,"index.html",namesInCities)
	if(err !=nil){
		log.Fatal(err)
	}

}
