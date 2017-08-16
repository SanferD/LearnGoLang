package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Once upon a $foo there was a $foo$foo with a $fo$foo and $fooo."
	fmt.Println(input)
	fmt.Println(expand(input, swap))
}

func swap(s string) string {
	if s == "foo" {
		return "ohno!"
	}

	return s
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}