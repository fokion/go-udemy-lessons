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
	err = tmpl.Execute(os.Stdout,nil)
	if err != nil{
		log.Fatal(err)
	}
}
