package main

import "fmt"

const arrayLength = 10

func main() {
	x := [arrayLength]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Printf("Before reverse %v\n", x)
	x = reverse(x)
	fmt.Printf("After reverse %v\n", x)
}

func reverse(array [arrayLength]int) [arrayLength]int {
	for i, j := 0, arrayLength-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
