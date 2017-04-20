package main

import (
	"text/template"
	"log"
	"os"
)

func main(){
	tmpl , err := template.ParseGlob("templates/*.tmpl")
	if(err != nil){
		log.Fatal(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout,"file2.tmpl",nil)
	if(err != nil){
		log.Fatal(err)
	}
}
