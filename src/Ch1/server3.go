package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var lock sync.Mutex

func main() {
	http.HandleFunc("/", handler)

	port := 8000
	fmt.Println("listening on port " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "%q\n", r)
}

