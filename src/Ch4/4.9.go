package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordFreqCounter := make(map[string]int)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open error: %v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordFreqCounter[word]++
	}
	if err := input.Err; err != nil {
		fmt.Fprintf(os.Stderr, "error scanning file: %v\n", err)
	}
	for word, count := range wordFreqCounter {
		fmt.Printf("%20s: %d\n", word, count)
	}
	file.Close()
}