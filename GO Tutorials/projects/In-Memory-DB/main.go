package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(connection)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// intructions
	io.WriteString(conn, "\nIN-Memory Database\n\n")

	// read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	// infinite looping to read text
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln) // slice  the string
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "Expected value")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID Cmd")
		}
	}
}
