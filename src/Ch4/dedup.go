package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	setOfNewLines := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	performDupOnInput(setOfNewLines, input)
	checkForError(input)
}

func performDupOnInput(setOfLines map[string]bool, input *bufio.Scanner) {
	for hasInput(input) {
		checkInput(setOfLines, input)
	}
}

func hasInput(input *bufio.Scanner) bool {
	return input.Scan()
}

func checkInput(setOfLines map[string]bool, input *bufio.Scanner) {
	line := getLine(input)
	performDupOnLine(setOfLines, line)
}

func getLine(input *bufio.Scanner) string {
	return input.Text()
}

func performDupOnLine(setOfLines map[string]bool, line string) {
	if isLineNew(setOfLines, line) {
		recordNewLine(setOfLines, line)
	}
}

func isLineNew(setOfLines map[string]bool, line string) bool {
	return !setOfLines[line]
}

func recordNewLine(setOfLines map[string]bool, line string) {
	setOfLines[line] = true
	fmt.Println(line)
}


func checkForError(input *bufio.Scanner) {
	if scannerHasError(input) {
		printErrorInScanner(input)
	}
}

func scannerHasError(input *bufio.Scanner) bool {
	return input.Err() != nil
}

func printErrorInScanner(input *bufio.Scanner) {
	err := input.Err()
	fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
}