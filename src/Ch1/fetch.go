package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", error)
			continue
		}
		body, error := ioutil.ReadAll(response.Body)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, error)
			os.Exit(1)
		}
		response.Body.Close()
		fmt.Printf("%s", body)
	}
}