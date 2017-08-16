package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	bytesReader := bytes.NewReader(p)
	scanner := bufio.NewScanner(bytesReader)
	scanner.Split(bufio.ScanWords)
	
	var numberOfWords int
	for scanner.Scan() {
		numberOfWords++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error scanning format string %s: %v", string(p), err)
	}
	
	*wc += WordCounter(numberOfWords)
	return numberOfWords, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	bytesReader := bytes.NewReader(p)
	scanner := bufio.NewScanner(bytesReader)
	scanner.Split(bufio.ScanLines)

	var numberOfLines int
	for scanner.Scan() {
		numberOfLines++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error scanning format string %s: %v", string(p), err)
	}

	*lc += LineCounter(numberOfLines)
	return numberOfLines, nil
}

func main() {
	var wc WordCounter
	_, err := fmt.Fprintf(&wc, "hello goodmorning %s %s !\n How is it going?\n", "Sanfer", "Dsouza")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Println(wc)

	var lc LineCounter
	_, err = fmt.Fprintf(&lc, "hello goodmorning %s %s !\n How is it going?\n", "Sanfer", "Dsouza")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Println(lc)
}