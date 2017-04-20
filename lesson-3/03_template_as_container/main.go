package main

import (
	"text/template"
	"log"
	"os"
)

func main(){
	tmpl , err := template.ParseFiles("file1")
	if err != nil{
		log.Fatal(err)
	}
	tmpl , err = tmpl.ParseFiles("file2","file3")
	if err != nil{
		log.Fatal(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout,"file2",nil)
	if err != nil{
		log.Fatal(err)
	}
}
