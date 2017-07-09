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
type Item struct{
	Name string
	Completed bool
}
type Todo struct {
	Name string
	Items[]Item
}
func main(){
	myTodo := Todo{
		Name: "Testing",
		Items: []Item{
			{Name:"First Element",Completed:false},
			{Name:"Second element",Completed:true},
		},
	}
	err:= tmpl.ExecuteTemplate(os.Stdout,"index.html",myTodo)
	if(err != nil){
		log.Fatal(err)
	}
}