package main

import "fmt"

func main() {
	myarray := [9]int{0, 1,2,3,4,5,6,7,8};
	myslice := myarray[2:5]
	fmt.Println(myslice)

	fmt.Println("before reverse ", myarray)
	reverse(myarray[3:6])
	fmt.Println("after reverse ", myarray)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}