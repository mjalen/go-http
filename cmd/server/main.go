package main

import (
	"net"
	"flag"
	"log"
	"fmt"

	"go-http/http/route"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "The port number for the server to listen to.")
	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port)) 
	if err != nil {
		log.Fatalf("ERR: %s", err)
	}
	log.Printf("Server listening on port %v.", port)

	router := route.Router{
		"/": func(conn net.Conn) {
			fmt.Fprintf(conn, "Hello there :)\r\n")	
		},
		"/test": func(conn net.Conn) {
			fmt.Fprintf(conn, "This is the test route. :)\r\n")
		},
	}

	for {
		conn, err := ln.Accept()
		// once HTTP/1.0 is implemented replace with a 500?
		if err != nil { 
			log.Printf("ERR (Accept): %s", err)
			continue
		}

		go func() {
			req := make([]byte, 1024)
			_, err := conn.Read(req)
			if err != nil {
				log.Printf("ERR (Read): %s", err)
			}

			log.Printf("Got request: %s", req)
			router.Process(req, conn)
			conn.Close() // update to send info back, and dynamicly close based on connection requirements.
		}()
	}
}
