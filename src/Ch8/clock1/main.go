package main

import (
	"net"
	"time"
	"fmt"
	"os"
	"log"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}
		handleConnection(con)
	}
}

func handleConnection(con net.Conn) {
	defer con.Close()
	fmt.Fprintf(con, "%s\n", time.Now().Format(time.Kitchen))
}