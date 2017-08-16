package main

import "fmt"

func main() {
	x := 20
	y := 20.34
	z := "abc"
	fmt.Printf("%T %[1]v\n", x)
	fmt.Printf("%T %[1]v\n", y)
	fmt.Printf("%T %[1]v\n", z)
}