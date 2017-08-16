package main

import "fmt"

func main() {
	str := "abcdefghijkl"

	byteArray := convertStringToByteArray(str)
	printByteArray("Before", byteArray)
	byteArray = reverse(byteArray)
	printByteArray("After", byteArray)
}

func reverse(array []byte) []byte {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func convertByteArrayToString(byteArray []byte) string {
	return string(byteArray)
}

func convertStringToByteArray(str string) []byte {
	return []byte(str)
}

func printByteArray(desc string, byteArray []byte) {
	stringFromBytes := convertByteArrayToString(byteArray)
	printIt(desc, stringFromBytes)
}

func printIt(desc string, s string) {
	fmt.Printf("%s\t%v\n", desc, s)
}
