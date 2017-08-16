package main

import (
	"bufio"
	"fmt"
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
			fmt.Printf("Listener accept error in main: %s", err.Error())
		}
		go handleConnection(connection)
	}
}

func handleConnection(c net.Conn) {
	inputCh := make(chan string)
	go readInput(c, inputCh)

	for {
		ticker := time.NewTicker(10 * time.Second)
		select {
		case <- ticker.C:
			fmt.Fprintln(c, "Timeout")
			c.Close()
			return
		case s := <- inputCh:
			go echo(c, s, time.Second)
		}
		ticker.Stop()
	}
}

func readInput(c net.Conn, ch chan<-string) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		ch <- input.Text()
	}
}

func echo(c net.Conn, s string, t time.Duration) {
	fmt.Fprintf(c, "\t%s\n", strings.ToUpper(s))
	time.Sleep(t)
	fmt.Fprintf(c, "\t%s\n", s)
	time.Sleep(t)
	fmt.Fprintf(c, "\t%s\n", strings.ToLower(s))
}
