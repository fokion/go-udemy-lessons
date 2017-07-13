package main

import (
	"net/http"
	"fmt"
)

type RandomHandler int
func (m RandomHandler) ServeHTTP( w http.ResponseWriter , r *http.Request){
	fmt.Fprintln(w,"anything you want")
}
func main(){
	var handler RandomHandler
	http.ListenAndServe(":80",handler)
}
