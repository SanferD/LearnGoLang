package main

import "fmt"

func main() {
	fmt.Println(max(0))
	fmt.Println(max(0,1))
	fmt.Println(max(1,0))
	fmt.Println(max(0, 10, 21, 2, 85, 30, 4, 36, 13, 14, 8))
}

func max(x int, args ...int) int {
	if len(args) == 0 {
		return x
	}

	max := x
	for _, value := range args {
		if value > max {
			max = value
		}
	}

	return max
}