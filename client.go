package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	defer conn.Close()

	response, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	response = strings.TrimSpace(response)
	if response != "OK" {
		log.Fatalf("Unexpected response: got %q, want %q", response, "OK")
	}

	fmt.Print("Received expected response:", response)
}
