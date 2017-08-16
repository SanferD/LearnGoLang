package main

import (
	"fmt"
	"Ch6/geometry"
)

func main() {
	path := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(path.Distance())
}