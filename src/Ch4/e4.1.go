package main

import (
	"crypto/sha512"
	"fmt"
	"os"
	"Ch2/popcount"
)

const sizeOfSha = 512

func main() {
	validateCmdInput()
	hash1, hash2 := computeHashValues()
	count := countDifferentBits(&hash1, &hash2)
	printTheCount(count)
}

func validateCmdInput() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}
}

func computeHashValues() ([sizeOfSha/8]uint8, [sizeOfSha/8]uint8) {
	hash1 := sha512.Sum512([]byte(os.Args[1]))
	hash2 := sha512.Sum512([]byte(os.Args[2]))
	return hash1, hash2
}

func countDifferentBits(A *[sizeOfSha/8]uint8, B *[sizeOfSha/8]uint8) int {
	count := 0
	for i, _ := range A {
		count += doTheCount(A[i], B[i])
	}
	return count
}

func doTheCount(one uint8, two uint8) int {
	diff := getDisjointUnion(one, two)
	return countBitsThatAreOne(diff)
}


func getDisjointUnion(first uint8, second uint8) uint8 {
	return (first - second ) | (second - first)
}

func countBitsThatAreOne(word uint8) int {
	return popcount.PopCount(uint64(word))
}

func printTheCount(count int) {
	fmt.Printf("Num of different bits: %d\n", count)
}