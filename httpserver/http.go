package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func exec(conn net.Conn) {
	defer conn.Close()
	var buf = make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			return
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("Client> %s\n", buf)
	fmt.Printf("Clinet> Size %d\n", n)

	n, err = conn.Write(buf[:n])
	if err != nil {
		log.Fatal(err)
	}
}
