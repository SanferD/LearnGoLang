// xkcd max is 1797

package main

import (
	"fmt"
	"os"
	"net/http"
	"encoding/json"
)

type Info struct {
	URL string
	Transcript string
}

func main() {
	var info Info
	info.URL = "https://xkcd.com/571/info.0.json"
	resp, err := http.Get(info.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http Get error: %v\n", err)
		resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Search query failed %s\n", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decoder error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(info.URL)
	fmt.Println(info.Transcript)

	resp.Body.Close()
}