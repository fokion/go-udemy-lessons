package main

import (
	"html/template"
	"os"
	"log"
	"math"
)

var tmpl *template.Template
func doubleIt(num float64) float64{
	return 2*num
}
func powerIt(num float64) float64{
	return math.Pow(num,2)
}
var templateFunctions = template.FuncMap{
	"dbl":doubleIt,
	"pow":powerIt,
}
func init(){
	tmpl = template.Must(template.New("index.html").Funcs(templateFunctions).ParseFiles("views/index.html"))
}

func main(){
	var num float64 = 3
	err:= tmpl.ExecuteTemplate(os.Stdout,"index.html",num)
	if(err != nil){
		log.Fatal(err)
	}
}