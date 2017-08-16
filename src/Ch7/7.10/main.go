package main

import (
	"Ch7/7.10/stringholder"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	for _, arg := range os.Args[1:] {
		s := stringholder.MakeStringHolder(arg)
		check := IsPallindrome(s)
		fmt.Printf("%s\t%v\n", arg, check)
	} 
}

func IsPallindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}