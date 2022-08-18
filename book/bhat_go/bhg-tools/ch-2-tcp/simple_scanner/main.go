package main

import (
	"fmt"
	"net"
)

// to see an open port client sends a -> syn
// if the response from server is a `syn-ack` the port is open (3 way handshake)
// if the response from server is a `rst` the port is closed
// if the response from server is non-existent/left on read (timeout) it is blocked by firewall

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("connection successful")
	} else {
		fmt.Println("connection no-good")
	}
}

// successful: <nil>
// no connection: dial tcp <address>: connect: connection refused
