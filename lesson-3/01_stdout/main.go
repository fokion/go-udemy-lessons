package main

import (
	"text/template"
	"log"
	"os"
)

func main(){
	tmpl , err := template.ParseFiles("tpl-gohtml")
	if err != nil{
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout,nil)
	if err != nil{
		log.Fatal(err)
	}
}
