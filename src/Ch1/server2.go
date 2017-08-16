package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var lock sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)

	port := 8000
	fmt.Println("listening on port " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	count++
	lock.Unlock()
	fmt.Fprintf(w, "Get: %s\n", r.URL.Path)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	lock.Unlock()
}