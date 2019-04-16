package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	// Scan() returns a bool
	for scanner.Scan() {
		// default is a line; use e.g., Split(bufio.ScanWords) to change that behavior
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard you say: %s\n", line)
	}
	defer conn.Close()

	// the connection will timeout and break out of the scanner loop
	// the following will execute as a result
	fmt.Println("Code got here.")
}

// 1. go run main.go
// 2. telnet localhost:8080
