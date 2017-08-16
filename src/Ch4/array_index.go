package main

import "fmt"

func main() {
	const (
		USD = iota
		PND
		SAR
		IR
	)

	values := [...]string{USD: "theUSD", PND: "thePND", SAR: "theSAR", IR: "theIR"}
	fmt.Println(values[USD])
	fmt.Println(values[PND])
	fmt.Println(values[SAR])
	fmt.Println(values[IR])
}