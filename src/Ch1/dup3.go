package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

func main() {
	counter := make(map[string]int)
	lineToFileNameMap := make(map[string]string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counter[line]++
			lineToFileNameMap[line] = filename
		}
	}

	for line, counts := range counter {
		if counts > 1 {
			var filename string = lineToFileNameMap[line]
			fmt.Printf("%v\t%d\t%v\n", filename, counts, line)
		}
	}
}