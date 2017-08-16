package main

import (
	"fmt"
	"unicode"
)

func main() {
	mystring := "This    string has   a   lot of  unnecessary    spaces     ... ."
	byteArray := []byte(mystring)
	
	printByteArray("Before", byteArray)
	byteArray = removeAdjacentDuplicates(byteArray)
	printByteArray("After", byteArray)
}

func removeAdjacentDuplicates(s []byte) []byte {
	for i := 0; i<len(s)-1; i++ {
		for areAdjacentValuesSpaces(s, i) {
			s = removeAdjacentDuplicate(s, i)
		}
	}
	return s
}

func areAdjacentValuesSpaces(s []byte, i int) bool {
	first, second := rune(s[i]), rune(s[i+1])
	return	unicode.IsSpace(first) &&
		 	unicode.IsSpace(second)
}

func removeAdjacentDuplicate(s []byte, i int) []byte {
	return append(s[:i+1], s[i+2:]...)
}

func printByteArray(desc string, byteArray []byte) {
	stringFromBytes := convertByteArrayToString(byteArray)
	printIt(desc, stringFromBytes)
}

func printIt(desc string, s string) {
	fmt.Printf("%s\t%v\n", desc, s)
}

func convertByteArrayToString(byteArray []byte) string {
	return string(byteArray)
}