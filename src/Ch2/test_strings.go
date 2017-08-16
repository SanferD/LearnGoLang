package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	slash := strings.LastIndex(arg, "/")
	arg = arg[slash+1:]
	dot := strings.LastIndex(arg, ".")
	if dot != -1 {
		arg = arg[:dot]
	}
	fmt.Println(arg)
}

