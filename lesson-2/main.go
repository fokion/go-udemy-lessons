package main

import (
	"os"
	"log"
	"io"
	"strings"
)

func main(){
nf, err := os.Create("hello.txt")
	if(err != nil){
		log.Fatal("Error creating file",err)
	}
	defer nf.Close()
	io.Copy(nf,strings.NewReader("Fokion"))
}
