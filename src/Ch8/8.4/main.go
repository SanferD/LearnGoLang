package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
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
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 2*time.Second, wg)
	}
	if err := input.Err(); err != nil {
		fmt.Printf("Scanner error in handleConnection: %s", err.Error())
	}
	go func() {
		wg.Wait()
		tcpConn := c.(*net.TCPConn)
		tcpConn.CloseWrite()
	}()
}

func echo(c net.Conn, s string, t time.Duration, wg sync.WaitGroup) {
	fmt.Fprintf(c, "\t%s\n", strings.ToUpper(s))
	time.Sleep(t)
	fmt.Fprintf(c, "\t%s\n", s)
	time.Sleep(t)
	fmt.Fprintf(c, "\t%s\n", strings.ToLower(s))
	wg.Done()
}
