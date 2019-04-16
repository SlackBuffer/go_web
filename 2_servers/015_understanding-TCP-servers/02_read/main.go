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
	}
	defer conn.Close()

	// the code never get here (visit using browser)
	// the scanner doesn't if the input is over and just keeps scanning (infinite loop)
	fmt.Println("Code got here.")
}

// 1. go run main.go
// 2. visit localhost:8080 (using firefox)
// 3. console will print http request
