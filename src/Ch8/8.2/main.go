package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
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
	input := bufio.NewScanner(c)
	for {
		if input.Scan() {
			cmd := input.Text()
			handleCommand(c, strings.TrimSpace(cmd))
		}
	}
}

func handleCommand(c net.Conn, cmd string) {
	values := strings.SplitN(cmd, " ", 2)
	if values == nil {
		fmt.Fprintf(c, "Invalid command: '%s'\n", cmd);
		return
	}

	cmd = values[0]
	var proc *exec.Cmd
	if cmd == "cd" {
		if len(values) == 1 {
			return
		}
		if err := os.Chdir(values[1]); err != nil {
			fmt.Fprintf(c, err.Error())
			return
		}
		fmt.Fprintf(c, "Changed directory to: %s\n", values[1])
	} else {
		if len(values) == 1 {
			proc = exec.Command(cmd)
		} else {
			args := strings.Split(values[1], " ");
			proc = exec.Command(cmd, args[1:]...)
		}
		outputInBytes, err := proc.Output()
		if err != nil {
			fmt.Fprintf(c, err.Error())
			return
		}
		output := string(outputInBytes)
		fmt.Fprintf(c, "%s\n", output)
	}
}
