package main

import (
	"context"
	"log"
)

func main() {

	info := &ServerInfo{
		Addr: "localhost",
		Port: "8080",
		Proto: "tcp",
	}

	srv, err := ServerInit(info, context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := srv.Sock.Accept()
		if err != nil {
			continue
		}
		go exec(conn)
	}
}
