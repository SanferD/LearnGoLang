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

	go broadcast()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err.Error())
			continue
		}
		go handleConnection(connection)
	}
}

type client struct {
	C chan <- string
	Name string
}
var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)
type tableOfClients map[client]bool

func broadcast() {
	clients := make(tableOfClients)
	for {
		select {
		case client := <-entering:
			registerNewClient(clients, client)			
		case client := <-leaving:
			removeClient(clients, client)
		case msg := <-messages:
			for client := range clients {
				client.C<-msg
			}
		}
	}
}

func registerNewClient(table tableOfClients, c client) {
	table[c] = true
	sendWelcomeMessage(table, c)
}

func sendWelcomeMessage(table tableOfClients, c client) {
	c.C <- "Welcome to chat!"
	if len(table) == 1 {
		c.C <- "You're the only one!"
	} else {
		c.C <- "The current clients are: "
		for client := range table {
			c.C <- client.Name 
		}
	}
}

func removeClient(clients tableOfClients, client client) {
	go func(name string) {messages <- name + " is leaving."}(client.Name)
	delete(clients, client)
	close(client.C)
}

func handleConnection(c net.Conn) {
	msgs := make(chan string)
	who := getIdentity(c)
	me := client{msgs, who}
	entering <- me

	go sendMessageToBroadcast(c, me)
	receiveMessages(c, msgs)
}

const timeoutTime time.Duration = 10*time.Minute
func sendMessageToBroadcast(c net.Conn, me client) {
	defer func() {
		leaving <- me
		c.Close()
	}()

	ticker := time.NewTicker(timeoutTime)
	go func() {
		<-ticker.C
		c.Close()
	}()

	fmt.Fprintln(c, "You are " + me.Name)
	input := bufio.NewScanner(c)
	for input.Scan() {
		msg := me.Name + ": " + input.Text()
		messages<-msg
		ticker.Stop()
		ticker = time.NewTicker(timeoutTime)
	}
}

func getIdentity(c net.Conn) string {
	remote := c.RemoteAddr().String()
	remoteSplit := strings.Split(remote, ":")
	return remoteSplit[len(remoteSplit)-1]
}

func receiveMessages(c net.Conn, mymessages <- chan string) {
	for msg := range mymessages {
		fmt.Fprintln(c, msg)
	}
}
