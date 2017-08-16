package main

import "fmt"

func main() {
	x := []int{0,1,2,3,4,5,6}
	i := 2

	fmt.Printf("Before removing i=%d\t%v\n", i, x)
	x = removePreserveOrder(x, i)
	fmt.Printf("After removing (preserve order) i=%d\t%v\n", i, x)

	x = removeDontPreserveOrder(x, 2)
	fmt.Printf("After removing (dont preserve order) i=%d\t%v\n", i, x)
}

func removePreserveOrder(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func removeDontPreserveOrder(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}