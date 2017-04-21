package main

import (
"html/template"
"os"
"strings"
)

var tmpl *template.Template


type Person struct {
	Name string
	Surname string
}
type Car struct {
	Name string
	Date string
}
type DB struct {
	People []Person
	Cars []Car
}
var templateFunctions = template.FuncMap{
	"threeLetters":firstThreeLetters,
	"lowerCase":strings.ToLower,
}
func firstThreeLetters(s string) string{
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}
func init(){
	tmpl = template.Must(template.New("index.html").Funcs(templateFunctions).ParseFiles("views/index.html"))
}
func main()  {
	var people = []Person{}
	p1 := Person{"Fokion","Sotiropoulos"}
	people = append(people,p1)
	people = append(people,Person{"Natalia", "Sotiropoulou"})
	var cars = []Car{}
	cars = append(cars,Car{"Toyota","1989"})
	cars = append(cars,Car{"Prius","2012"})
	data := DB{people,cars}
	tmpl.ExecuteTemplate(os.Stdout,"index.html",data)
}
