package main

import (
	"fmt"
	"os"
	"strconv"

	"Ch2/popcount"
)

func main() {
	var x uint64 = 3
	if len(os.Args) > 1 {
		y, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		x = uint64(y)
	}
	count := popcount.PopCount(x)
	fmt.Println(count)
}