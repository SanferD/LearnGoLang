package main

import (
	"fmt"
	"time"
)

func main() {
	abort := make(chan struct{})

	select {
	case <-time.After(10*time.Second):
		fmt.Println("ready to launch")
	case <-abort:
		fmt.Println("launch aborted")
		return
	}
	fmt.Println("LIFT OFF !")
}
