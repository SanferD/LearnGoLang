package main 

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<-string) {
	start := time.Now()
	response, error := http.Get(url)
	if error != nil {
		ch <- fmt.Sprint(error)
		return
	}

	nbytes, error := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close()
	if error != nil {
		ch <- fmt.Sprint("while reading %s: %v", url, error)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs,%d,%s", secs, nbytes, url)
}

// ./fetchall http://facebook.com http://gopl.io http://google.com http://umn.edu http://amazon.com http://netflix.com http://stackoverflow.com | sed \$d | sort -t"," -k2n,2 | tr , "\t"
