package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		response, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", error)
			os.Exit(1)
		}

		_, error = io.Copy(os.Stdout, response.Body)
		if error != nil {
			fmt.Fprintf(os.Stderr, "error: %s: %v\n", url, error)
			os.Exit(1)
		}
	}
} 