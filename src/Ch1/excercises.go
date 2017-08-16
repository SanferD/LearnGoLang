package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println("E1.1")
	fmt.Println(strings.Join(os.Args, " "))

	fmt.Println("E1.2")
	for index, arg := range os.Args {
		fmt.Println(index, " ", arg)
	}
}