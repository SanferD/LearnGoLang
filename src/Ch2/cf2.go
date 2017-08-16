package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"Ch2/tempconv"
	"Ch2/weightconv"
	"Ch2/heightconv"
)

func main() {
	if hasArguments() {
		printValuesForArguments()
	}	else {
		printValuesForStdin()
	}
}

func hasArguments() bool {
	return len(os.Args) > 1
}

func printValuesForArguments() {
		for _, arg := range os.Args[1:] {
			printValuesForArgument(arg)
		}
}

func printValuesForArgument(arg string) {
	number := getFloat64FromString(arg)
	printValues(number)
}

func printValuesForStdin() {
		input := bufio.NewScanner(os.Stdin)
		for hasInput(input) {
			printValuesForInput(input)
		}
		checkErrorFatally(input.Err())
}

func hasInput(input *bufio.Scanner) bool {
	return input.Scan()
}

func printValuesForInput(input *bufio.Scanner) {
	number := getFloat64FromString(input.Text())
	printValues(number)
}

func getFloat64FromString(str string) float64 {
	number, err := strconv.ParseFloat(str, 64)
	checkErrorFatally(err)
	return number
}

func checkErrorFatally(err error) {
	if err != nil {
		log.Fatal("%v\n", err)
	}
}

func printValues(number float64) {
	printHeight(number)
	printWeight(number)
	printTemperature(number)
}

func printHeight(number float64) {
	ft, m := heightconv.Foot(number), heightconv.Meter(number)
	fmt.Printf("height: \n%s\n%s\n\n", heightconv.MToF(m), heightconv.FToM(ft))
}

func printWeight(number float64) {
	p, k := weightconv.Pound(number), weightconv.Kilogram(number)
	fmt.Printf("weight: \n%s\n%s\n\n", weightconv.PToK(p), weightconv.KToP(k))
}

func printTemperature(number float64) {
	f, c := tempconv.Farenheit(number), tempconv.Celsius(number)
	fmt.Printf("temperature: \n%s\n%s\n\n", tempconv.CToF(c), tempconv.FToC(f))
}