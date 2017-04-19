package main

import "fmt"

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
}