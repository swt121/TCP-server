package main

import (
	"fmt"
	"log"
	"net"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("OK\n"))

	if err != nil {
		fmt.Println("Error writing to connection")
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Error listening on %s: %v", ":8080", err)
	}

	defer listener.Close()
	fmt.Println("Listening on " + listener.Addr().String())

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		fmt.Println("Connected with", conn.RemoteAddr().String())
		go handleRequest(conn)
	}
}
