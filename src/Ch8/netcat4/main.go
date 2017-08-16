package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	connection := getConnection()
	go mustCopy(os.Stdout, connection)
	mustCopy(connection, os.Stdin)
}

func getConnection() net.Conn {
	connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err.Error())
	}
	return connection
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
