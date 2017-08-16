package main

import "fmt"

type S struct {
	a, b int
}

func main() {
	var s1, s2 S
	s1.a = 1;
	s1.b = 2;
	s2 = s1;
	s1.a = 3;
	fmt.Println(s2.a)
}
