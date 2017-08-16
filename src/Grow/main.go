package main

import "fmt"

type S struct {
	Size int
}

type K struct {
	S S
}

func main() {
	var k K
	k.S.Size += 10
	fmt.Println(k.S.Size)
}
