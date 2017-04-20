package main

import (
	"text/template"
	"log"
	"os"
)

var tmpl *template.Template
/*Initializes the app and it is running at the beginning and only once*/
func init(){
	//the must is doing the error checking internally
	tmpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}
func main(){
	err := tmpl.ExecuteTemplate(os.Stdout,"file2.tmpl",nil)
	if(err != nil){
		log.Fatal(err)
	}
}
