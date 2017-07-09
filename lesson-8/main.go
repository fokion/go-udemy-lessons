package main

import (
	"net"
	"log"

	"fmt"

	"bufio"
	"time"
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
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(30*time.Second))
	fmt.Fprintf(conn,"write something : \n")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn,"I got : "+ln+"\n")
	}
	defer conn.Close()
	fmt.Println("connection closed")
}


