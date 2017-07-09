package main

import (
	"html/template"
	"os"
	"log"
)

var tmpl *template.Template

func init(){
	tmpl, _ = template.ParseFiles("views/index.html","views/item.html")
}

func main(){
	var num float64 = 3
	err:= tmpl.ExecuteTemplate(os.Stdout,"index.html",num)
	if(err != nil){
		log.Fatal(err)
	}
}