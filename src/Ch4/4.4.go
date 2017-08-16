package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5}
	i := 2

	printIt("Before rotate", x, i)
	x = rotate(x, i)
	printIt("After rotate", x, i)
}

func rotate(s []int, i int) []int {
	return append(s[i:], s[:i]...)
}

func printIt(desc string, s []int, i int) {
	fmt.Printf("%s i=%d\t%v\n", desc, i, s)
}