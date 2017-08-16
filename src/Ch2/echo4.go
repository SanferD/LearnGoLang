package main

import (
	"flag"
	"fmt"
	"strings"
)

var toPrintNewLine = flag.Bool("n", false, "print new line")
var wordSeparator = flag.String("s", " ", "separator used between words")

func main() {
	flag.Parse()
	printWithoutNewline()
	if hasNewLine() {
		printNewLine()
	}
}

func printWithoutNewline() {
	fmt.Print(strings.Join(flag.Args(), *wordSeparator))
}

func hasNewLine() bool {
	return *toPrintNewLine
}

func printNewLine() {
	fmt.Println()
}
