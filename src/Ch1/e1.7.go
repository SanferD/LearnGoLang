package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v\n", error)
			os.Exit(1)
		}

		_, error = io.Copy(os.Stdout, response.Body)
		if error != nil {
			fmt.Fprintf(os.Stderr, "copying body error %s: %v\n", error)
			os.Exit(1)
		}
	}
}