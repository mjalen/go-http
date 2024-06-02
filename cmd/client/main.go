package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"

	"go-http/http/syntax"
)

func main() {
	var port int
	var path string
	flag.IntVar(&port, "port", 8080, "Port to send request/message.")
	flag.StringVar(&path, "path", "/", "Path of the desired route.")
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	defer conn.Close()
	if err != nil {
		log.Fatalf("ERR (Dial): %s", err)
	}

	req := fmt.Sprintf("GET %s\r\n", path)
	ok := syntax.Validate("Simple-Request", []byte(req)) 
	if !ok {
		log.Fatalf("ERR: Malformed request.")
	}
	log.Printf("Request: %s", req)
	fmt.Fprintf(conn, "%s", req)

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalf("ERR (ReadString): %s", err)
	}

	log.Printf("Got response: %s", status)
}
