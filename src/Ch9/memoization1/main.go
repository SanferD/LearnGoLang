package main

import (
	"bufio"
	"fmt"
	"Ch9/memoization1/memo"
	"os"
)

func CountNumberOfWords(name string) (interface{}, error) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("CountNumberOfWords couldn't open file %s: %s", name, err.Error())
	}
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	var count int
	for input.Scan() {
		count++
	}
	if err:=input.Err(); err != nil {
		return nil, fmt.Errorf("CountNumberOfWords error with scanner on file %s: %s", name, err.Error())
	}

	return count, nil
}

func main() {
	m := memo.New(CountNumberOfWords)
	fileNames := []string{ "main.go", "memo/memo.go", }
	for _, name := range fileNames {
		count, err := m.Get(name)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%s: %d\n", name, count)
		}
	}
}