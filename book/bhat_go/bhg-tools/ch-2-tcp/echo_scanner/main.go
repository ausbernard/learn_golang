package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	// create a buffer to store the recieved data
	b := make([]byte, 512)
	for {
		// recieve data via conn.Read into buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("[-]client is down.")
			break
		}
		if err != nil {
			log.Println("[-]unexpected error")
			break
		} else {
			log.Printf("[+]Received %d bytes: %s\n", size, string(b))
		}
		// data is received and stored in buffer
		// send data via conn.Write from buffer
		log.Println("writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// bind app to TCP port 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("[-]Unable to bind to port")
	}
	log.Println("[+]Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
}
