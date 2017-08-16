package main

import (
	"fmt"
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), time.Second)
	}
	c.Close()
}

func echo(c net.Conn, s string, t time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(t)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(t)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}
