package main

import (
	"log"
	"net"
)

func main() {
	srv, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			continue
		}
		go exec(conn)
	}
}
