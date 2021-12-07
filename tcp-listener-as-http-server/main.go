package main

import "net"
import "fmt"
import "strconv"

// https://stackoverflow.com/questions/33293293/simple-web-server-implementation-golang
func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Cannot listen on port 8080: ", err)
		return
	}
	// run loop forever (or until ctrl-c)
	for {
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error when accepting connection: ", err)
			return
		}

		// no request processing, simple response
		conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
		conn.Write([]byte("Content-Type: text/plain; charset=UTF-8\r\n"))
		resp := "Hello World!"
		conn.Write([]byte("Content-Length: " + strconv.Itoa(len(resp)) + "\r\n\r\n"))
		conn.Write([]byte(resp + "\n"))
	}
}
