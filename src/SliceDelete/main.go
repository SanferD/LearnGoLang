package main

import "fmt"

type S struct {
	elmts []int
}

func main() {
	var obj S
	obj.elmts = []int{1, 2, 3, 4, 5}
	fmt.Println(obj.elmts)
	reverse(obj.elmts)
	fmt.Println(obj.elmts)
}

func reverse(s []int) {
	for i:=0; i!=len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func remove(s []int, i int) {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	s = s[:len(s)-1]
}
