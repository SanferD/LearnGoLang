package main

import (
	"Ch7/tempconv"
	"flag"
	"fmt"
)

var celsius = tempconv.CelsiusFlag("temperature", tempconv.Celsius(20.0), "<temp><C, F, or K>. eq -temperature 20F")

func main() {
	flag.Parse()
	fmt.Printf("Celsius: %v\n", *celsius)
}