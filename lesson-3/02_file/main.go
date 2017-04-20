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
	file , err := os.Create("index.html")
	if err !=nil{
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.Execute(file,nil)
	if err != nil{
		log.Fatal(err)
	}
}
