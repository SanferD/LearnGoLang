package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := make(map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, counter)
	}	else {
		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countLines(file, counter)
			file.Close()
		}
	}

	for line, counts := range counter {
		if counts > 1 {
			fmt.Printf("%d\t%s\n", counts, line)
		}
	}
}

func countLines(file *os.File, counter map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counter[input.Text()]++
	}
}