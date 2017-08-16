package main

import (
	"bufio"
	"Ch7/7.5/specialreader"
	"io"
	"strings"
	"fmt"
)

func LimitReader(r io.Reader, n int64) io.Reader {
	newReader := specialreader.Create(r, int(n))
	return newReader
}

func main() {
	mystring := "Long ago in a distant land I Akku the shape shifting master of darkness"
	strReader := strings.NewReader(mystring)
	reader := LimitReader(strReader, 25)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}