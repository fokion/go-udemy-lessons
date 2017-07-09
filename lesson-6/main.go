package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

func main(){
	li , err := net.Listen("tcp",":8080")
	if err != nil{
		log.Panic(err)
	}
	defer li.Close()

	for{
		conn, err := li.Accept()
		if err != nil{
			log.Panic(err)
		}
		io.WriteString(conn,"Hello from TCP server")
		fmt.Fprintf(conn,"How is your day?")
		conn.Close()
	}
}


