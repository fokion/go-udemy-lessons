package main

import "fmt"

type person struct {
	firstName string
	lastName string
}
type agent struct {
	person
	licenceToKill bool
}
func (p person) speak(){
	fmt.Println(p.firstName,`says ,"Good morning"`)
}
func (sa agent) speak(){
	fmt.Println(sa.firstName , sa.lastName ,`says ,"Shaken , not stirred"`)
}
//polymorphism
type human interface {
	speak()
}
func saySomething(h human){
	h.speak()
}
func main() {
	x := 7
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	//this is a slice of ints..
	xi := []int{1,2,3,4,5,6}
	fmt.Println(xi)
	//this is a map of string with int
	//notice the trailing comma in the
	//last element.
	m:=map[string]int{
		"Fokion":28,
		"Alexandra":26,
		"Natalia":25,
	}
	fmt.Println(m)
	p1 := person{"Miss","Moneypenny"}
	p1.speak()
	ag1 := agent{person{"James","Bond"},true}
	ag1.speak()
	//get the speak method for the person
	ag1.person.speak()
	saySomething(p1)
	saySomething(ag1)
}
