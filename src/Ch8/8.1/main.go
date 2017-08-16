package main

import (
	"fmt"
	"flag"
	"log"
	"net"
	"time"
)

var ip = flag.Int("port", 8000, "The port number for localhost")
const delay = time.Second

func main() {
	flag.Parse()
	host := fmt.Sprintf("localhost:%d", *ip)
	fmt.Println(host)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	for {
		fmt.Fprintf(c, "The time is: %s\n", time.Now().Format(time.Stamp))
		time.Sleep(delay)
	}
}