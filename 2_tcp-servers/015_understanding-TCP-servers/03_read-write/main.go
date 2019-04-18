package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
	scanner := bufio.NewScanner(conn)
	// Scan() returns a bool
	for scanner.Scan() {
		// default is a line; use e.g., Split(bufio.ScanWords) to change that behavior
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard you say: %s\n", line)
	}
	defer conn.Close()

	// the code never get here
	// the scanner doesn't if the input is over and just keeps scanning (infinite loop)
	fmt.Println("Code got here.")
}

// 1. go run main.go
// 2. telnet localhost:8080
