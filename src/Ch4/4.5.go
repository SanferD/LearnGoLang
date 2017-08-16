package main

import "fmt"

func main() {
	values := []string{"apple", "oranges", "oranges", "oranges", "oranges", "pears", "bannanas", "bannanas", "teapot"}
	
	printIt("Before", values)
	values = removeAdjacentDuplicates(values)
	printIt("After", values)
}

func removeAdjacentDuplicates(s []string) []string {
	for i := 0; i<len(s)-1; i++ {
		for areAdjacentValuesEqual(s, i) {
			s = removeAdjacentDuplicate(s, i)
		}
	}
	return s
}

func areAdjacentValuesEqual(s []string, i int) bool {
	return s[i] == s[i+1]
}

func removeAdjacentDuplicate(s []string, i int) []string {
	return append(s[:i+1], s[i+2:]...)
}

func printIt(desc string, s []string) {
	fmt.Printf("%s\t%v\n", desc, s)
}