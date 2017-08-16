package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err.Error())
	}

	done := make(chan struct{})
	go func(dst io.Writer, src io.Reader) {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatal(err)
		}
		done<-struct{}{}
	}(os.Stdout, connection)
	mustCopy(connection, os.Stdin)
	connection.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
