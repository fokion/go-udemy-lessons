package main

import (
	"html/template"
	"time"
	"os"
	"log"
)

var tmpl *template.Template
func monthDayYear(t time.Time) string{
	//based on 01/02 03:04:05PM '06-0700
	//01 month , 02 day 03 h , 04 min , 05 sec ...
	return t.Format("01-02-2006")
}
var templateFunctions = template.FuncMap{
	"MDY":monthDayYear,
}
func init(){
	tmpl = template.Must(template.New("index.html").Funcs(templateFunctions).ParseFiles("views/index.html"))
}

func main(){
	err:= tmpl.ExecuteTemplate(os.Stdout,"index.html",time.Now())
	if(err != nil){
		log.Fatal(err)
	}
}