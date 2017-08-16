package main

import "fmt"

func main() {
	fmt.Println(whatever())
}

func whatever() (x int) {
	defer func() {
		if p := recover(); p != nil {
			x = 7
		}
	}()
	panic("yelp")
}