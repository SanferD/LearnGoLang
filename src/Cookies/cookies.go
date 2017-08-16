package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", saveCookie)
	http.HandleFunc("/print", sendCookie)
	log.Fatal( http.ListenAndServe("localhost:9000", nil) )
}

func saveCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "ssid", Value: "1234"}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Set cookie\n")
}

func sendCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("ssid")
	fmt.Fprintf(w, "Got cookie %v\n", cookie)
}

