// cf converts its numerical arguments from celsius to farenheit
package main

import (
	"fmt"
	"os"
	"strconv"

	"Ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf error: %v\n", err)
			os.Exit(1)
		}
		f, c := tempconv.Farenheit(t), tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}